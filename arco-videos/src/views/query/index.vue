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
  <a-grid :cols="6" :rowGap="5" :colGap="5"  class="grid-demo-grid">
    <a-grid-item class="demo-item" :span="1" v-for="x in contentData" v-bind:key="x" > 
        <a-card class="card-demo" >
    <template #cover>
      <div :style="{height: '260px'}">
        <img
          :style="{ width: '100%',height: 'auto'}"
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
    <a-link :href="`/info?id=${UrlParse(x.url)}`">进入详情</a-link>
    <a-link @click="star(UrlParse(x.url),x.title,x.thumb)">
      收藏
      <template #icon>
        <icon-font :type="exists(UrlParse(x.url))" :size="15"/>
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
const starFiled = ref(0)
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_x0593fl85v.js' });
const starList = ref([])
axios.get(`/video/starList`).then(response => {
  if (response.code === 200){
    response.data.forEach(element => {
      starList.value.push(element.tvId)
    });
    console.log(starList)
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

function log() {
  if (!proxy.$cookies.get("urk")){
    visible.value +=1
  }else{
    axios.get(`/video/query?key=${queryKey.value}`).then(response => {
      if (response.code === 7001){
        proxy.$message.error(response.msg)
        return
      }
        contentData.value = response
    }).catch(error => {
       // 请求失败处理
      console.log(error);
    });
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
      proxy.$message.success("收藏成功")
      axios.get(`/video/query?key=${queryKey.value}`).then(queryResponse => {
      if (queryResponse.code === 7001){
        proxy.$message.error(queryResponse.msg)
        return
      }
      contentData.value = queryResponse
    }).catch(error => {
       // 请求失败处理
      console.log(error);
    });
    }).catch(error => {
       // 请求失败处理
      console.log(error);
    });
}
</script>