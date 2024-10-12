<template>
    <a-grid :cols="10" class="grid-demo-grid" :collapsed="collapsed">
        <br>
        <a-grid-item class="demo-item" :span="18"></a-grid-item>
        <a-grid-item class="demo-item" :span="2" :offset="1"> <a-link href="/query" style="text-decoration:underline;">&lt;&lt;返回</a-link></a-grid-item>
    </a-grid>
    <a-grid :cols="8" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
        <a-grid-item class="demo-item" :span="8"></a-grid-item>
        <a-grid-item class="demo-item" :offset="1" :span="6">
            <div id="xgPlayerWrap"></div>
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
import { ref,onMounted } from 'vue';
import Player,{Events} from 'xgplayer'
import 'xgplayer/dist/index.min.css'
import { useRoute } from 'vue-router'
import HlsPlugin from 'xgplayer-hls'
import { Icon } from '@arco-design/web-vue';
import axios from 'axios';
import { conf } from "./conf";

const contentData = ref([])
const playStruct = ref({
    teleplay:"",
    key:"",
    name:""
})
let player = null // 实例//
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_of4nm2nht7e.js' });
const init = (playUrl,startTime2) => {
    player = new Player({
        url:playUrl,
        ...conf,
        startTime:startTime2,
        plugins: [HlsPlugin]
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
const getInfo = (url)=>{
    axios.get(`/video/getInfoV2?url=${url.split("/")[2]}`).then(response => {
        contentData.value = response
        Object.keys(response).forEach(function(key,index) {
            if (response[key].info.play === "1"){
                defaultIndex.value = index
            }
            
        })
        console.log(defaultIndex)
    }).catch(error => {
        // 请求失败处理
        console.log(error);
    });
}
// 播放视频
function playerD(value,key,name,startTime){
    console.log(startTime)
    init(value,startTime)
    const teleplay =route.query.url.split("/")[2]
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
     getInfo(route.query.url)
     
})
</script>