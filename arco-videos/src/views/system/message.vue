<template>
  <div class="message-container">
    <div class="message-header">
      <h2 class="message-title">留言板</h2>
      <a-button type="primary" @click="showMessageModal">
        发表留言
      </a-button>
    </div>

    <div class="message-list">
      <a-list :data="messages" :bordered="false">
        <template #item="{ item }">
          <a-list-item class="message-item">
            <div class="message-content">
              <div class="message-user">
                <a-avatar :size="40">
                  <img v-if="item.avatar" :src="item.avatar" />
                  <icon-user v-else />
                </a-avatar>
                <span class="username">{{ item.username }}</span>
                <span class="time">{{ item.createTime }}</span>
              </div>
              <div class="message-text">{{ item.content }}</div>
              <div class="message-actions">
                <a-space>
                  <a-button type="text" size="small" @click="handleReply(item)">
                    <icon-message /> 回复
                  </a-button>
                  <a-button v-if="item.userId === currentUserId" type="text" status="danger" size="small" @click="handleDelete(item)">
                    <icon-delete /> 删除
                  </a-button>
                </a-space>
              </div>
              <!-- 回复列表 -->
              <div v-if="item.replies && item.replies.length > 0" class="reply-list">
                <div v-for="reply in item.replies" :key="reply.id" class="reply-item">
                  <a-avatar :size="24">
                    <img v-if="reply.avatar" :src="reply.avatar" />
                    <icon-user v-else />
                  </a-avatar>
                  <span class="username">{{ reply.username }}</span>
                  <span class="reply-content">{{ reply.content }}</span>
                  <span class="time">{{ reply.createTime }}</span>
                </div>
              </div>
            </div>
          </a-list-item>
        </template>
      </a-list>
    </div>

    <!-- 发表留言对话框 -->
    <a-modal
      v-model:visible="messageModalVisible"
      :title="isReply ? '回复留言' : '发表留言'"
      @ok="handleSubmit"
      @cancel="handleCancel"
      :mask-closable="false"
      :footer-class="'modal-footer'"
    >
      <a-form :model="form" ref="formRef">
        <a-form-item field="content" :rules="[{ required: true, message: '请输入留言内容' }]">
          <a-textarea
            v-model="form.content"
            :placeholder="isReply ? '请输入回复内容' : '请输入留言内容'"
            :auto-size="{ minRows: 3, maxRows: 8 }"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import axios from 'axios'

const messages = ref([]);
const messageModalVisible = ref(false);
const isReply = ref(false);
const currentUserId = ref('');
const replyToMessage = ref(null);

const form = ref({
  content: '',
  parentId: null
});

const formRef = ref(null);

// 获取留言列表
function fetchMessages (){
  axios.get(`/message/list`).then(response => {
    console.log(response)
    if (response.code === 200) {
      messages.value = response.data;
    }
  }).catch(error => {
    console.error('获取留言失败:', error);
    Message.error('获取留言失败');
  });
};

// 显示发表留言对话框
const showMessageModal = () => {
  isReply.value = false;
  form.value = { content: '', parentId: null };
  messageModalVisible.value = true;
};

// 处理回复
const handleReply = (message) => {
  isReply.value = true;
  replyToMessage.value = message;
  form.value = { content: '', parentId: message.id };
  messageModalVisible.value = true;
};

// 处理删除
const handleDelete = async (message) => {
  try {
    const response = await axios.post('/message/delete', { id: message.id });
    if (response.code === 200) {
      Message.success('删除成功');
      await fetchMessages();
    } else {
      Message.error(response.msg || '删除失败');
    }
  } catch (error) {
    console.error('删除失败:', error);
    Message.error('删除失败');
  }
};

// 提交留言或回复
const handleSubmit = async () => {
  try {
    await formRef.value.validate();
    const url = isReply.value ? '/message/reply' : '/message/add';
    const response = await axios.post(url, form.value);
    
    if (response.code === 200) {
      Message.success(isReply.value ? '回复成功' : '发表成功');
      messageModalVisible.value = false;
      await fetchMessages();
    } else {
      Message.error(response.msg || (isReply.value ? '回复失败' : '发表失败'));
    }
  } catch (error) {
    console.error('提交失败:', error);
    Message.error('提交失败');
  }
};

const handleCancel = () => {
  messageModalVisible.value = false;
  form.value = { content: '', parentId: null };
};

onMounted(() => {
  fetchMessages();
});
</script>

<style scoped>
.message-title {
  color: rgb(var(--red-6));
  font-size: 28px;
  font-weight: bold;
  text-shadow: 3px 3px 6px rgba(0, 0, 0, 0.15);
  position: relative;
  padding-bottom: 8px;
  transition: transform 0.3s ease;
}

.message-title:hover {
  transform: translateY(-2px);
}

.message-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 60px;
  height: 3px;
  background-color: rgb(var(--red-6));
  border-radius: 2px;
  transition: width 0.3s ease;
}

.message-title:hover::after {
  width: 100px;
}

.message-container {
  padding: 20px;
  margin: 80px auto;
  min-height: calc(100vh - 160px);
  overflow-y: auto;
  background: linear-gradient(to bottom, #ffffff, var(--color-bg-1));
  border-left: 4px solid rgba(var(--primary-6), 0.2);
  border-right: 4px solid rgba(var(--primary-6), 0.2);
  box-shadow: 
    -8px 0 15px -5px rgba(var(--primary-6), 0.1),
    8px 0 15px -5px rgba(var(--primary-6), 0.1);
  position: relative;
  margin-left: 20px;
  margin-right: 20px;
  border-radius: 16px;
  overflow: hidden;
  max-width: 1200px;
}

/* 修改容器的渐变边框效果 */
.message-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: -4px;
  right: -4px;
  height: 100px;
  background: linear-gradient(to bottom,
    rgba(var(--primary-6), 0.15),
    transparent
  );
  pointer-events: none;
  border-radius: 16px 16px 0 0;
}

.message-container::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: -4px;
  right: -4px;
  height: 100px;
  background: linear-gradient(to top,
    rgba(var(--primary-6), 0.15),
    transparent
  );
  pointer-events: none;
  border-radius: 0 0 16px 16px;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
  position: relative;
  z-index: 1;
  padding: 20px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(var(--primary-6), 0.1);
}

.message-item {
  padding: 20px;
  border-radius: 12px;
  background: linear-gradient(to bottom, #ffffff, var(--color-bg-2));
  margin-bottom: 24px;
  box-shadow: 
    0 4px 6px -1px rgba(0, 0, 0, 0.1),
    0 2px 4px -1px rgba(0, 0, 0, 0.06),
    inset 0 2px 4px rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(var(--primary-6), 0.1);
  border-left: 4px solid rgba(var(--primary-6), 0.3);
  transition: all 0.3s ease;
  transform: translateY(0);
  position: relative;
  overflow: hidden;
  backdrop-filter: blur(8px);
}

.message-item:hover {
  transform: translateY(-2px);
  box-shadow: 
    0 10px 15px -3px rgba(var(--primary-6), 0.15),
    0 4px 6px -2px rgba(var(--primary-6), 0.1),
    inset 0 2px 4px rgba(255, 255, 255, 0.3);
  border-color: rgba(var(--primary-6), 0.3);
  border-left-color: rgba(var(--primary-6), 0.6);
  background: linear-gradient(to bottom, #ffffff, var(--color-bg-1));
}

.message-content {
  position: relative;
  z-index: 1;
}

.message-user {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
  position: relative;
  z-index: 1;
}

.username {
  margin-left: 8px;
  font-weight: 500;
  color: var(--color-text-1);
  transition: color 0.3s ease;
}

.username:hover {
  color: rgb(var(--primary-6));
}

.time {
  margin-left: 16px;
  color: var(--color-text-3);
  font-size: 12px;
}

.message-text {
  margin: 8px 0;
  color: var(--color-text-1);
  line-height: 1.6;
  padding: 8px;
  border-radius: 8px;
  background-color: var(--color-bg-1);
  transition: background-color 0.3s ease;
}

.message-text:hover {
  background-color: var(--color-bg-3);
}

.message-actions {
  margin-top: 8px;
  opacity: 0.6;
  transition: opacity 0.3s ease;
}

.message-item:hover .message-actions {
  opacity: 1;
}

.reply-list {
  margin-top: 16px;
  padding-left: 48px;
  border-left: 2px solid rgba(var(--primary-6), 0.2);
  transition: border-color 0.3s ease;
}

.message-item:hover .reply-list {
  border-left-color: rgba(var(--primary-6), 0.6);
}

.reply-item {
  display: flex;
  align-items: center;
  padding: 8px 12px;
  font-size: 14px;
  opacity: 0.9;
  transition: all 0.3s ease;
  border-radius: 10px;
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0.8), var(--color-bg-2));
  box-shadow: 
    0 2px 4px -1px rgba(0, 0, 0, 0.06),
    inset 0 1px 2px rgba(255, 255, 255, 0.1);
  margin-bottom: 8px;
}

.reply-item:hover {
  opacity: 1;
  transform: translateX(4px);
  background: linear-gradient(to bottom, #ffffff, var(--color-bg-1));
  box-shadow: 
    0 4px 6px -1px rgba(var(--primary-6), 0.1),
    inset 0 1px 2px rgba(255, 255, 255, 0.2);
}

.reply-content {
  margin: 0 8px;
  color: var(--color-text-1);
  background-color: var(--color-bg-1);
  padding: 4px 8px;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

.reply-content:hover {
  background-color: var(--color-bg-3);
}

/* 添加列表动画 */
.message-list-enter-active,
.message-list-leave-active {
  transition: all 0.5s ease;
}

.message-list-enter-from,
.message-list-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

/* 添加按钮动画 */
.arco-btn {
  transition: all 0.3s ease !important;
}

.arco-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(var(--primary-6), 0.2);
}

/* 添加头像动画 */
.arco-avatar {
  transition: all 0.3s ease;
}

.arco-avatar:hover {
  transform: scale(1.1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 移动端适配 */
@media screen and (max-width: 600px) {
  .message-container {
    padding: 16px;
    margin: 60px 8px;
    border-left-width: 2px;
    border-right-width: 2px;
    border-radius: 12px;
  }

  .message-header {
    padding: 16px;
    margin-bottom: 24px;
  }

  .message-item {
    padding: 16px;
    margin-bottom: 20px;
  }

  .reply-list {
    padding-left: 24px;
  }

  /* 移动端对话框额外调整 */
  :deep(.arco-modal) {
    width: 95vw !important;
    max-height: 80vh;
  }

  :deep(.arco-modal-header) {
    padding: 12px 16px;
  }

  :deep(.arco-modal-body) {
    padding: 12px;
  }

  :deep(.arco-form-item) {
    margin-bottom: 12px;
  }

  :deep(.arco-textarea-wrapper) {
    font-size: 14px;
  }

  :deep(.arco-btn) {
    font-size: 14px;
    height: 32px;
    line-height: 32px;
  }

  /* 移动端优化动画效果 */
  .message-item:hover {
    transform: none;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  }
  
  .reply-item:hover {
    transform: none;
  }
}

/* 移除列表项的边框 */
:deep(.arco-list-item) {
  border: none;
  padding: 0;
}

:deep(.arco-list) {
  border: none;
}

/* 移动端对话框样式优化 */
:deep(.arco-modal) {
  max-width: 90vw;
  margin: 0 auto;
  top: 50% !important;
  transform: translateY(-50%) !important;
}

:deep(.arco-modal-body) {
  max-height: 60vh;
  overflow-y: auto;
  padding: 16px;
}

:deep(.arco-textarea-wrapper) {
  width: 100%;
  max-height: 40vh;
  overflow-y: auto;
}

:deep(.arco-modal-footer) {
  padding: 12px 16px;
  border-top: 1px solid var(--color-border);
}

:deep(.arco-btn) {
  padding: 4px 16px;
}
</style>
