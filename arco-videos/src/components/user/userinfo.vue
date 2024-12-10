<template>
    <a-modal visible="true" title="用户设置" ok-text="修改" @cancel="handleCancel" @before-ok="handleBeforeOk">
        <a-form :model="form">
          <a-form-item field="name" label="用户名">
            <a-input v-model="form.name" />
          </a-form-item>
          <!-- <a-form-item field="post" label="头像">
            <a-space direction="vertical" :style="{ width: '100%' }">
                <a-upload
                  action="/"
                  :fileList="file ? [file] : []"
                  :show-file-list="false"
                  @change="onChange"
                  @progress="onProgress"
                >
                  <template #upload-button>
                    <div
                      :class="`arco-upload-list-item${
                        file && file.status === 'error' ? ' arco-upload-list-item-error' : ''
                      }`"
                    >
                      <div
                        class="arco-upload-list-picture custom-upload-avatar"
                        v-if="file && file.url"
                      >
                        <img :src="file.url" />
                        <div class="arco-upload-list-picture-mask">
                          <IconEdit />
                        </div>
                        <a-progress
                          v-if="file.status === 'uploading' && file.percent < 100"
                          :percent="file.percent"
                          type="circle"
                          size="mini"
                          :style="{
                            position: 'absolute',
                            left: '50%',
                            top: '50%',
                            transform: 'translateX(-50%) translateY(-50%)',
                          }"
                        />
                      </div>
                      <div class="arco-upload-picture-card" v-else>
                        <div class="arco-upload-picture-card-text">
                          <IconPlus />
                          <div style="margin-top: 10px; font-weight: 600">Upload</div>
                        </div>
                      </div>
                    </div>
                  </template>
                </a-upload>
              </a-space>
          </a-form-item> -->
        </a-form>
      </a-modal>
</template>

<script setup>
import {ref,watch} from 'vue'
import { IconEdit, IconPlus } from '@arco-design/web-vue/es/icon';

const props = defineProps({
  visibleParent: Boolean,
});
const visible = ref(true)
const file = ref();
watch(() => props.visibleParent,(n,o)=>{
  visible.value = true
})
const onChange = (_, currentFile) => {
      file.value = {
        ...currentFile,
        // url: URL.createObjectURL(currentFile.file),
      };
    };
const onProgress = (currentFile) => {
    file.value = currentFile;
};
function handleCancel(){
    visible.value = false
}
function handleBeforeOk(){
    visible.value = false
}
</script>