<template>
     
  <a-grid :cols="3" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
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
    <!-- <a-link :href="`http://xdm530.com${x.url}`">进入详情</a-link> -->
    <a-link :href="`/info?url=${x.url}`">进入详情</a-link>
  </a-card>
    </a-grid-item>
    
  </a-grid>
  
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
import { ref } from 'vue';
import axios from 'axios'

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
</script>