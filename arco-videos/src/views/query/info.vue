<template>
    <a-grid :cols="10" class="grid-demo-grid" :collapsed="collapsed">
        <br>
        <a-grid-item class="demo-item" :span="18"></a-grid-item>
        <a-grid-item class="demo-item" :span="2" :offset="1"></a-grid-item>
    </a-grid>
    <a-grid :cols="8" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
        <a-grid-item class="demo-item" :span="8"></a-grid-item>
        <a-grid-item class="demo-item" :offset="1" :span="6">
            <div id="xgPlayerWrap"></div>
        </a-grid-item>
    </a-grid>
    <br/>
    <a-grid  :cols="48" class="grid-demo-grid" :collapsed="collapsed">
        <a-grid-item class="demo-item"  :offset="1" :span="1">
                <a-button @click="download()">
                    <template #icon>
                    <icon-font type="icon-xiazai3" :size="26"/>
                </template> 
            </a-button>
         </a-grid-item>
    </a-grid>
    <br/>
    <a-tabs type="rounded" size="medium"  animation="true">
        <a-tab-pane v-for="(value,key,index) in contentData" v-bind:key="index" >
            <template #title>
                <icon-font type="icon-bofang" :size="15" v-if="value.info.play == 1"/>{{index+1}}
              </template>
            <a-grid :cols="8" :colGap="2" :rowGap="20" class="grid-demo-grid" :collapsed="collapsed">
                <a-grid-item class="demo-item" :span="1" style="display:block;margin:0 auto" v-for="(item) in value.list" v-bind:key="item">
                    <a-button type="dashed" status="success" @click="playerD(item.url,key,item.name,item.startTime)">
                        <template #icon v-if="item.play == 1">
                            <icon-font type="icon-bofang" :size="15"/>
                        </template>
                        {{item.name}}
                    </a-button>
                </a-grid-item>
            </a-grid>
        </a-tab-pane>
      </a-tabs>
</template>

<style>
#xgPlayerWrap { flex: auto; }
#xgPlayerWrap video { width: 100%; }
</style>

<script setup>
import { ref,onMounted,getCurrentInstance } from 'vue';
import Player,{Events} from 'xgplayer'
import 'xgplayer/dist/index.min.css'
import { useRoute } from 'vue-router'
import HlsPlugin from 'xgplayer-hls'
import Mobile from 'xgplayer/es/plugins/mobile'
import { Icon } from '@arco-design/web-vue';
import axios from 'axios';
import { conf } from "./conf";

const {proxy} = getCurrentInstance()
const downloadUrl = ref("")
const contentData = ref([])
const playStruct = ref({
    teleplay:"",
    key:"",
    name:""
})

let player = null // 实例//
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_ptsrwy5a1fm.js' });
const init = (playUrl,startTime2) => {
    player = new Player({
        url:playUrl,
        ...conf,
        startTime:startTime2,
        plugins: [HlsPlugin,Mobile]
    });
    player.on(Events.SEEKED, (ev) => {
        axios.get(`/video/playRecord?teleplay=${playStruct.value.teleplay}&index=${playStruct.value.key}&name=${playStruct.value.name}&startTime=${ev.currentTime}`).then(response => {
        }).catch(error => {});
    })
    player.on(Events.TIME_UPDATE, (ev) => {
        if (ev.currentTime % 60 === 0){
        axios.get(`/video/playRecord?teleplay=${playStruct.value.teleplay}&index=${playStruct.value.key}&name=${playStruct.value.name}&startTime=${ev.currentTime}`).then(response => {
        }).catch(error => {});
        }
    })
    // 等各种监听事件
}

const route=useRoute()
const defaultIndex = ref(0)
// 获取视频信息
const getInfo = (id)=>{
    axios.get(`/video/getInfoV2?url=${id}`).then(response => {
        contentData.value = response
        Object.keys(response).forEach(function(key,index) {
            if (response[key].info.play === "1"){
                defaultIndex.value = index
            }

        })
    }).catch(error => {
        // 请求失败处理
        console.log(error);
    });
}

function download(){
    if (downloadUrl.value === ""){
        proxy.$message.error("请选择要下载的内容")
        return
    }
    axios.get(`/video/download?url=${downloadUrl.value}`).then(response => {
        if(response.code !== 200){
            proxy.$message.error(response.msg)
        }
    }).catch(error => {});
}
// 播放视频
function playerD(value,key,name,startTime){
    downloadUrl.value = value
    init(value,startTime)
    const teleplay =route.query.id
    playStruct.value.teleplay = teleplay
    playStruct.value.key = key
    playStruct.value.name = name
    axios.get(`/video/playRecord?teleplay=${teleplay}&index=${key}&name=${name}`).then(response => {
        }).catch(error => {});
    axios.get(`/video/getInfoV2?url=${teleplay}`).then(responses => {
        contentData.value = responses
        }).catch(error => {
    });
    
}

onMounted(() => {
    init("",0) 
    getInfo(route.query.id)
     
})
</script>