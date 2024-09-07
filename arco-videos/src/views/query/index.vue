<template>
    <a></a>
  <a-grid :cols="3" :colGap="12" :rowGap="30" class="grid-demo-grid" :collapsed="collapsed">
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1"></a-grid-item>
    <a-grid-item class="demo-item" :offset="1">
        <a-input v-model="queryKey"  placeholder="请输入剧名，开启奇妙之旅" class="input-rounded " @press-enter="log"/>
    </a-grid-item>
  </a-grid>
<br/>
  <a-grid :cols="6" :rowGap="30" :colGap="10"  class="grid-demo-grid">
    <a-grid-item class="demo-item" :span="1" @v-for="x in data"> 
        <a-card hoverable>
    <template #cover>
      <div >
        <img
          :style="{ width: '100%', transform: 'translateY(-20px)' }"
          alt="dessert"
          src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a20012a2d4d5b9db43dfc6a01fe508c0.png~tplv-uwbnlip3yd-webp.webp"
        />
      </div>
    </template>
    <a-card-meta title="Card Title">
      <!-- <template #description>
        Card content <br />
        Card content
      </template> -->
    </a-card-meta>
  </a-card>
    </a-grid-item>
    
  </a-grid>
</template>
<style lang="css">
.input-rounded{
    font-size: x-large;
        border-radius:14px;
        text-align: center;
        background-color:white;
        position: relative;
        border: 3px solid #dddddd;
        padding: 10px;
}
</style>

<script setup>
const queryKey = defineModel()
function log() {
    console.log(queryKey)
}
const data = [{
      key: '1',
      name: 'Jane Doe',
      salary: 23000,
      address: '32 Park Road, London',
      email: 'jane.doe@example.com'
    }, {
      key: '2',
      name: 'Alisa Ross',
      salary: 25000,
      address: '35 Park Road, London',
      email: 'alisa.ross@example.com'
    }, {
      key: '3',
      name: 'Kevin Sandra',
      salary: 22000,
      address: '31 Park Road, London',
      email: 'kevin.sandra@example.com'
    }, {
      key: '4',
      name: 'Ed Hellen',
      salary: 17000,
      address: '42 Park Road, London',
      email: 'ed.hellen@example.com'
    }, {
      key: '5',
      name: 'William Smith',
      salary: 27000,
      address: '62 Park Road, London',
      email: 'william.smith@example.com'
    }]
</script>