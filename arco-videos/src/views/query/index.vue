<template>
  <a-spin :loading="loading" tip="搜索中..." style="width: 100%">
    <br/>
    <a-grid :cols="3" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
      <a-grid-item class="demo-item" :offset="1"></a-grid-item>
      <a-grid-item class="demo-item" :offset="1"></a-grid-item>
      <a-grid-item class="demo-item" :span="2">
        <div class="search-container">
          <a-input 
            v-model="queryKey"  
            placeholder="请输入剧名，开启奇妙之旅" 
            class="input-rounded" 
            @press-enter="log"
            :loading="loading"
          />
          <a-button 
            type="primary" 
            class="search-button"
            @click="log"
            :loading="loading"
          >
            <template #icon>
              <icon-search />
            </template>
            搜索
          </a-button>
        </div>
      </a-grid-item>
    </a-grid>
    <a-alert 
      class="alert-warning"
      type="error"
      closable>
      如果想看的电视、电影为红色边框请重新搜索
      <div class="feedback-tip">
      如有反馈请前往<router-link to="/message" class="message-link">留言板</router-link>
    </div>
    </a-alert>

  <br/>
    <a-collapse v-model:activeKey="activeKey" :default-active-key="1" style="width: 98%;margin-left: 1%;">
      <a-collapse-item header="收藏列表" key="1">
        <div class="scroll-container">
          <div class="scroll-wrapper">
            <a-grid :cols="7" :rowGap="12" :colGap="12" class="grid-demo-grid star-list-grid">
              <a-grid-item class="demo-item" :span="1" v-for="x in starListEle" :key="x"> 
                <a-card class="card-demo" size="small">
                  <template #cover>
                    <div class="card-cover">
                      <img
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
          </div>
        </div>
      </a-collapse-item>
    </a-collapse>

  <br>
    <div class="multi-scroll-container">
      <div class="multi-scroll-wrapper">
        <a-grid v-if="contentData && contentData.length > 0" :cols="6" :rowGap="12" :colGap="12" class="grid-demo-grid search-list-grid">
          <a-grid-item class="demo-item" :span="1" v-for="(item) in contentData" :key="item.title"> 
            <a-card 
              class="card-demo" 
              :class="{ 'highlight-card': UrlParse(item.url) < 100 }"
              size="small"
            >
              <template #cover>
                <div class="card-cover">
                  <img :src="item.thumb" />
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
      </div>
    </div>
    <login :visibleParent="visible"></login>
  </a-spin>
</template>
<style scoped>
/* 高亮卡片样式 */
.highlight-card {
  border: 2px solid #f53f3f !important;
}

/* 确保提示内容显示 */
:deep(.arco-alert) {
  margin: 0 auto;
  text-align: center;
  font-size: 16px;
}

@media screen and (max-width:600px){
  .grid-demo-grid {
    grid-template-columns: repeat(2, 1fr) !important;
    padding: 0 8px;
  }

  .card-demo {
    margin: 0 4px;
  }

  :deep(.card-demo img) {
    width: 100% !important;
    height: 100% !important;
    object-fit: cover;
  }

  :deep(.card-demo .arco-card-cover) {
    height: 160px !important;
  }

  :deep(.arco-card-meta-title) {
    font-size: 14px;
    line-height: 1.4;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }

  :deep(.arco-card-meta-description) {
    font-size: 12px;
    line-height: 1.4;
  }

  :deep(.arco-link) {
    font-size: 12px;
    margin-top: 4px;
    display: block;
  }

  :deep(.arco-card-body) {
    padding: 8px;
  }

  /* 搜索框样式调整 */
  .input-rounded {
    left: 0;
    width: 90%;
    margin: 0 5%;
  }

  .alert-warning {
    width: 90%;
    left: -6%;
  }
}

/* 默认样式 */
.grid-demo-grid {
  width: 100%;
  margin: 0 auto;
}

.card-demo {
  transition: transform 0.3s ease;
}

.card-cover {
  height: 260px;
  overflow: hidden;
}

.card-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

/* 卡片悬停效果 */
.card-demo:hover {
  transform: translateY(-4px);
}

.card-demo:hover .card-cover img {
  transform: scale(1.05);
}

#xgPlayerWrap { flex: auto; }
#xgPlayerWrap video { width: 100%; }

/* 添加横向滚动样式 */
.scroll-container {
  width: 100%;
  overflow: hidden;
  position: relative;
}

.scroll-wrapper {
  overflow-x: auto;
  overflow-y: hidden;
  white-space: nowrap;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE and Edge */
  padding: 10px 0;
}

.scroll-wrapper::-webkit-scrollbar {
  display: none; /* Chrome, Safari and Opera */
}

/* 移动端样式调整 */
@media screen and (max-width: 600px) {
  .scroll-wrapper .grid-demo-grid {
    display: flex !important;
    grid-template-columns: none !important;
  }

  .scroll-wrapper .demo-item {
    flex: 0 0 auto;
    width: 160px !important;
    margin-right: 12px;
  }

  .scroll-wrapper .card-demo {
    width: 100%;
    margin: 0;
  }

  .scroll-wrapper :deep(.card-demo .arco-card-cover) {
    height: 200px !important;
  }
}

/* 添加滑动指示器 */
.scroll-container::after {
  content: '';
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  width: 40px;
  background: linear-gradient(to right, transparent, var(--color-bg-2));
  pointer-events: none;
  opacity: 0.8;
}

.scroll-container::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 40px;
  background: linear-gradient(to left, transparent, var(--color-bg-2));
  pointer-events: none;
  opacity: 0.8;
  z-index: 1;
}

/* 优化卡片在滚动容器中的显示 */
.scroll-wrapper .card-demo {
  margin: 0;
  height: 100%;
}

.scroll-wrapper .demo-item {
  display: inline-block;
  vertical-align: top;
  white-space: normal;
}

/* 多方向滚动容器样式 */
.multi-scroll-container {
  width: 100%;
  height: calc(100vh - 280px); /* 增加高度，减少顶部预留空间 */
  overflow: hidden;
  position: relative;
  margin-top: 10px;
  border-radius: 8px;
  background: var(--color-bg-2);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.multi-scroll-wrapper {
  height: 100%;
  overflow: auto;
  padding: 16px;
  scrollbar-width: thin;
  scrollbar-color: var(--color-fill-3) transparent;
}

/* 自定义滚动条样式 */
.multi-scroll-wrapper::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.multi-scroll-wrapper::-webkit-scrollbar-track {
  background: transparent;
}

.multi-scroll-wrapper::-webkit-scrollbar-thumb {
  background-color: var(--color-fill-3);
  border-radius: 3px;
}

.multi-scroll-wrapper::-webkit-scrollbar-thumb:hover {
  background-color: var(--color-fill-4);
}

/* 移动端优化 */
@media screen and (max-width: 600px) {
  .multi-scroll-container {
    height: calc(100vh - 200px); /* 移动端也相应增加高度 */
    margin: 10px 0;
  }

  .multi-scroll-wrapper {
    padding: 8px;
  }

  .multi-scroll-wrapper .grid-demo-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr) !important;
    gap: 8px;
  }

  .multi-scroll-wrapper .demo-item {
    width: 100% !important;
    margin: 0;
  }

  .multi-scroll-wrapper .card-demo {
    margin: 0;
  }

  .multi-scroll-wrapper :deep(.card-demo .arco-card-cover) {
    height: 160px !important;
  }
}

/* 添加滚动阴影效果 */
.multi-scroll-container::after {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 20px;
  background: linear-gradient(to bottom, transparent, rgba(0, 0, 0, 0.05));
  pointer-events: none;
}

.multi-scroll-container::before {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  height: 20px;
  background: linear-gradient(to top, transparent, rgba(0, 0, 0, 0.05));
  pointer-events: none;
  z-index: 1;
}

/* 优化卡片在容器中的显示 */
.multi-scroll-wrapper .card-demo {
  height: 100%;
  margin: 0;
}

/* 添加滚动动画 */
.multi-scroll-wrapper {
  scroll-behavior: smooth;
}

/* 触摸设备优化 */
@media (hover: none) {
  .multi-scroll-wrapper {
    -webkit-overflow-scrolling: touch;
  }
}

/* 修改搜索容器样式 */
.search-container {
  display: flex;
  align-items: center;
  gap: 10px;
  position: relative;
  z-index: 100;
  padding: 10px;
  border-radius: 14px;
}

.input-rounded {
  flex: 1;
  left: 25%;
  font-size: x-large;
  border-radius: 14px;
  text-align: center;
  background-color: var(--color-fill-2);
  position: relative;
  border: 3px solid #dddddd;
  padding: 10px;
  z-index: 1;
}

.search-button {
  height: 40px;
  border-radius: 14px;
  font-size: 16px;
  position: relative;
  z-index: 2;
}

/* 移动端适配 */
@media screen and (max-width: 600px) {
  .search-container {
    flex-direction: column;
    gap: 8px;
    width: 100%;
    padding: 8px;
  }

  .input-rounded {
    left: 0;
    width: 90%;
    margin: 0 5%;
  }

  .search-button {
    width: 90%;
    margin: 0 5%;
  }
}

/* 添加全局loading样式 */
:deep(.arco-spin) {
  width: 100%;
  min-height: 100vh;
}

:deep(.arco-spin-loading) {
  background-color: rgba(255, 255, 255, 0.6);
}

:deep(.arco-spin-tip) {
  font-size: 16px;
  color: var(--color-text-2);
}

.alert-warning {
  width: 50%;
  margin: 8px auto;
  left: -6%;
  position: relative;
}

/* 添加反馈提示样式 */
.feedback-tip {
  text-align: center;
  margin: 8px auto;
  font-size: 14px;
  color: var(--color-text-2);
}

.message-link {
  color: rgb(var(--primary-6));
  text-decoration: none;
  transition: all 0.3s ease;
  padding: 2px 4px;
  border-radius: 4px;
}

.message-link:hover {
  color: rgb(var(--primary-7));
  background-color: rgba(var(--primary-6), 0.1);
  text-decoration: underline;
}

/* 移动端适配 */
@media screen and (max-width: 600px) {
  .feedback-tip {
    font-size: 12px;
    margin: 4px auto;
    width: 90%;
  }
}
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

.card-demo .arco-card-cover {
  overflow: hidden;
}

.card-demo .arco-card-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.card-demo:hover .arco-card-cover img {
  transform: scale(1.05);
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
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_1mxb9zlmcwj.js' });
const starList = ref([])
const starListEle = ref([])
const loading = ref(false);
const activeKey = ref(1);

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

async function queryVideo(){
  try {
    loading.value = true;
    const queryResponse = await axios.get(`/video/query?key=${queryKey.value}`);
    if (queryResponse.code === 7001){
      proxy.$message.error(queryResponse.msg);
      return;
    }
    console.log(queryResponse)
    if (!queryResponse || queryResponse === "    ") {
      contentData.value = [];
      console.log("查询结果为空")
      return;
    }
    contentData.value = queryResponse;
  } catch (error) {
    console.log(error);
    proxy.$message.error('搜索失败，请稍后重试');
    loading.value = false;
  } finally {
    loading.value = false;
  }
}

async function log() {
  if (!proxy.$cookies.get("urk")){
    visible.value += 1;
  } else {
    activeKey.value = []; // 收起收藏列表
    await queryVideo();
  }
}

// 添加刷新收藏列表的函数
function refreshStarList() {
  axios.get(`/video/starList`).then(res => {
    if (res.data) {
      starList.value = [];
      starListEle.value = res.data;
      res.data.forEach(element => {
        starList.value.push(element.tvId);
      });
    }
  }).catch(error => {
    console.log('获取收藏列表失败:', error);
    proxy.$message.error('刷新收藏列表失败');
  });
}

function star(pid,pname,purl){
  // 检查是否登录
  if (!proxy.$cookies.get("urk")){
    visible.value += 1
    return
  }

  axios.post("/video/star",{
    id:pid,
    name:pname,
    url:purl
  }).then(response => {
    console.log(response)
      // 处理响应数据
      if (response.code !== 200){
        proxy.$message.error(response?.msg || '收藏失败');
        return;
      }
      
      // 显示成功消息
      proxy.$message.success(response.msg);
      
      // 刷新收藏列表
      refreshStarList();
      
      // 刷新查询结果
      if (queryKey.value) {
        queryVideo();
      }
    }).catch(error => {
      console.log('收藏操作失败:', error);
      proxy.$message.error('收藏失败，请稍后重试');
    });
}

</script>