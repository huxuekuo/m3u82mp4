<template>
     <!-- <div id="xgPlayerWrap"></div> -->
  <a-grid :cols="3" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :span="2">
        <a-input v-model="queryKey"  placeholder="请输入剧名，开启奇妙之旅" class="input-rounded " @press-enter="log"/>
    </a-grid-item>
  </a-grid>
<br/>
  <a-grid :cols="6" :rowGap="5" :colGap="5"  class="grid-demo-grid">
    <a-grid-item class="demo-item" :span="1" v-for="x in contentData" v-bind:key="x" > 
        <a-card class="card-demo" >
    <template #cover>
      <div >
        <img
          :style="{ width: '100%', transform: 'translateY(-20px)' }"
          alt="dessert"
          :src="x.thumb"
        />
      </div>
    </template>
    <a-card-meta :title="x.title">
      <template #description>
          地区：{{x.area}}<br/>
          连载至：{{x.lianzaijs}}集<br/>
        </template>
    </a-card-meta>
    <a-link :href="`http://xdm530.com${x.url}`">进入详情</a-link>
  </a-card>
    </a-grid-item>
    
  </a-grid>
  
</template>
<style scoped>
#xgPlayerWrap { flex: auto; }
#xgPlayerWrap video { width: 100%; }
</style>
<style lang="css">
.input-rounded{
        left: 25%;
        font-size: x-large;
        border-radius:14px;
        text-align: center;
        background-color:white;
        position: relative;
        border: 3px solid #dddddd;
        padding: 10px;
}
.card-demo {
  /* width: 360px; */
  /* margin-left: 24px; */
  transition-property: all;
}
.card-demo:hover {
  transform: translateY(-4px);
}
</style>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue';
import Player,{Events} from 'xgplayer'
import 'xgplayer/dist/index.min.css'
import HlsPlugin from 'xgplayer-hls'
import axios from 'axios'
import { conf } from "./conf";



const queryKey = ref("")
const contentData = ref([])
function log() {
  axios.get(`/query?key=${queryKey.value}`).then(response => {
    contentData.value = response
    console.log(contentData)
          }).catch(error => {
              // 请求失败处理
              console.log(error);
          });
 
}

function cardClick(p){
  console.log(p)
}
let player = null // 实例

const init = () => {
    player = new Player({
        ...conf,
        plugins: [HlsPlugin]
    });
    player.on(Events.PLAY, (ev) => {
        console.log('-播放开始-', ev);
    })
    player.on(Events.PAUSE, (ev) => {
        console.log('-播放结束-', ev);
    })
    player.on('loadedmetadata', (ev) => {
        console.log('-媒体数据加载好了-', ev);
    })
    player.on(Events.SEEKED, (ev) => {
        console.log('-跳着播放-', ev);
    })
    // 等各种监听事件
}
onMounted(() => { init() })

</script>