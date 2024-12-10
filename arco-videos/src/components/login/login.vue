<template>
    <a-modal :visible="visible" title="登录提醒" @ok="handleOk" ok-text="登录并注册" @cancel="handleCancel">
        <a-form >
          <a-form-item  label="用户邮箱">
            <a-input v-model="ui.u"/>
          </a-form-item>
          <a-form-item  label="密码">
            <a-input-password v-model="ui.p"/>
          </a-form-item>
        </a-form>
      </a-modal>
</template>

<script setup>
import {watch,ref,getCurrentInstance} from 'vue'
import {useUserInfoStore} from '@/store'
import axios from 'axios'

const {proxy} = getCurrentInstance()
const ui = ref({
  u:"",
  p:"",
})

const userInfo = useUserInfoStore()

const props = defineProps({
  visibleParent: Boolean,
});

const visible = ref(false)
watch(() => props.visibleParent,(n,o)=>{
  visible.value = true
})
function handleOk(){
  axios.post(`/user/register`,ui.value).then(response => {
    if(response.code === 200){
      console.log(response.data.t)
      // 设置cookie
      proxy.$cookies.set('urk', response.data.t, '60d')
      visible.value = false
    }else{
      console.log(response.msg)
      proxy.$message.error(response.msg)
    }
  }).catch(error => {
  });
}
function handleCancel(){
  visible.value = false
}
</script>