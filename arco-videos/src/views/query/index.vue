<template>
  <br/>
  <a-grid :cols="24" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
    <a-grid-item class="demo-item" :offset="24"></a-grid-item>
  </a-grid>
  <br/>
  <a-grid :cols="3" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :span="2">
        <a-input v-model="queryKey"  placeholder="请输入剧名，开启奇妙之旅" class="input-rounded" @press-enter="log"/>
    </a-grid-item>
  </a-grid>
<br/>
  <a-collapse :default-active-key="1" style="width: 98%;margin-left: 1%;">
    <a-collapse-item header="收藏列表" key="1">
      <a-grid :cols="6" :rowGap="5" :colGap="5"  class="grid-demo-grid">
        <a-grid-item class="demo-item" :span="1" v-for="x in starListEle" :key="x" > 
          <a-card class="card-demo" >
      <template #cover>
        <div :style="{height: '260px'}">
          <img
            :style="{ width: '100%',height: 'auto'}"
            :src="x.url"
          />
        </div>
      </template>
      <a-card-meta :title="x.name">
      </a-card-meta>
      <a-link :href="`/info?id=${x.tvId}`">进入详情</a-link>
      <a-link @click="star(x.tvId,x.name,x.url)">
        收藏
        <template #icon>
          <icon-font :type="exists(x.tvId)" :size="15"/>
      </template>
      </a-link>
      </a-card>
      </a-grid-item>
      </a-grid>
    </a-collapse-item>
  </a-collapse>

<br>
  <a-grid :cols="6" :rowGap="5" :colGap="5"  class="grid-demo-grid">
    <a-grid-item class="demo-item" :span="1" v-for="(item) in contentData" :key="item.title"> 
        <a-card class="card-demo" >
    <template #cover>
      <div :style="{height: '260px'}">
        <img
          :style="{ width: '100%',height: 'auto'}"
          :src="item.thumb"
        />
      </div>
    </template>
    <a-card-meta :title="item.title">
      <template #description>
          地区：{{item.area}}<br/>
          连载至：{{item.lianzaijs}}集<br/>
        </template>
    </a-card-meta>
    <a-link :href="`/info?id=${UrlParse(item.url)}`">进入详情</a-link>
    <a-link @click="star(UrlParse(item.url),item.title,item.thumb)">
      收藏
      <template #icon>
        <icon-font :type="exists(UrlParse(item.url))" :size="15"/>
    </template>
    </a-link>
  </a-card>
    </a-grid-item>
  </a-grid>
  <login :visibleParent="visible"></login>
</template>
<style scoped>
@media screen and (max-width:600px){
  .card-demo {
    margin-left: 24px;
    margin-right: 24px;
  }
  .grid-demo-grid{
    grid-template-columns:repeat(1, minmax(0px, 1fr));
  }
}
#xgPlayerWrap { flex: auto; }
#xgPlayerWrap video { width: 100%; }
</style>
<style lang="css">
.input-rounded{
        left: 25%;
        font-size: x-large;
        border-radius:14px;
        text-align: center;
        background-color: var(--color-fill-2);
        position: relative;
        border: 3px solid #dddddd;
        padding: 10px;
}
input::input-placeholder{
	color:red;
}

.card-demo {
  transition-property: all;
}
.card-demo:hover {
  transform: translateY(-4px);
}

</style>

<script setup>
import { ref,getCurrentInstance } from 'vue';
import { Icon } from '@arco-design/web-vue';
import {UrlParse} from '@/utils/url'
import axios from 'axios'
import login from '../../components/login/login.vue'

const {proxy} = getCurrentInstance()
const visible = ref(0)
const queryKey = ref("")
const contentData = ref([])
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_x0593fl85v.js' });
const starList = ref([])
const starListEle = ref([])
axios.get(`/video/starList`).then(response => {
  if (response.code === 200){
      response.data.forEach(element => {
      starList.value.push(element.tvId)
      starListEle.value = response.data
    });
  }
}).catch(error => {
// 请求失败处理
console.log(error);
});

function exists(tvId){
  let b = "icon-dibudaohanglan-"
  starList.value.forEach(element => {
    if (element === tvId){
      b = "icon-shoucang-xingxing"
    }
    });
    return b
}
function queryVideo(){
  axios.get(`/video/query?key=${queryKey.value}`).then(queryResponse => {
      if (queryResponse.code === 7001){
        proxy.$message.error(queryResponse.msg)
        return
      }
      contentData.value = queryResponse
      console.log(contentData.value)
    }).catch(error => {
       // 请求失败处理
      console.log(error);
    });
}

function log() {
  if (!proxy.$cookies.get("urk")){
    visible.value +=1
  }else{
    queryVideo()
  }
}

function star(pid,pname,purl){
  axios.post("/video/star",{
    id:pid,
    name:pname,
    url:purl
  }).then(response => {
      if (response.code !== 200){
        proxy.$message.error(response.msg)
        return
      }
      proxy.$message.success(response.msg)
      queryVideo()
    }).catch(error => {
       // 请求失败处理
      console.log(error);
    });
}


</script>