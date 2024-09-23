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
    <a-tabs type="rounded" size="medium" >
        <a-tab-pane v-for="(value,key,index) in contentData" v-bind:key="value" >
            <template #title>
                <icon-font type="icon-bofang" :size="15" v-if="value.info.play == 1"/>{{index+1}}
              </template>
            <a-grid :cols="8" :colGap="2" :rowGap="20" class="grid-demo-grid" :collapsed="collapsed">
                <a-grid-item class="demo-item" :span="1" style="display:block;margin:0 auto" v-for="(item) in value.list" v-bind:key="item">
                    <a-button type="dashed" status="success" @click="playerD(item.url,key,item.name)">
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
let player = null // 实例//
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_of4nm2nht7e.js' });
const init = (playUrl) => {
    player = new Player({
        url:playUrl,
        ...conf,
        plugins: [HlsPlugin]
    });
    player.on(Events.SEEKED, (ev) => {
        console.log('-跳着播放-', ev);
    })
    player.on(Events.TIME_UPDATE, (ev) => {
        if (ev.currentTime % 300 == 0){

        }
    })
    // 等各种监听事件
}

const route=useRoute()
// 获取视频信息
const getInfo = (url)=>{
    axios.get(`/getInfoV2?url=${url.split("/")[2]}`).then(response => {
        contentData.value = response
          }).catch(error => {
              // 请求失败处理
              console.log(error);
          });
}
// 播放视频
function playerD(value,key,name){
    init(value) 
    
    const teleplay =route.query.url.split("/")[2]
   axios.get(`/playRecord?teleplay=${teleplay}&index=${key}&name=${name}`).then(response => {
    }).catch(error => {});
    axios.get(`/getInfoV2?url=${teleplay}`).then(responses => {
        contentData.value = responses
    }).catch(error => {
        console.log(error);
    });
}

onMounted(() => {
    init("") 
     getInfo(route.query.url)
     
})
</script>