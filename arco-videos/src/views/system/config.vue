<template>
    <div class="config-container">
      <div style="margin-bottom: 16px;">
        <a-button type="primary" @click="handleAdd">新增字典</a-button>
      </div>
      
      <a-table :columns="columns" :data="data" :pagination="false">
        <template #status="{ record }">
          <a-tag :color="record.state ? 'green' : 'red'">
            {{ record.state ? '启用' : '禁用' }}
          </a-tag>
        </template>
        <template #optional="{ record }">
          <a-space>
            <a-button type="primary" size="small" @click="handleEdit(record)">编辑</a-button>
            <a-button type="primary" status="danger" size="small" @click="handleDelete(record)">删除</a-button>
          </a-space>
        </template>
      </a-table>
      
      <a-modal v-model:visible="visible" @ok="handleOk" @cancel="handleCancel">
        <template #title>
          {{ editingRecord ? '编辑字典' : '新增字典' }}
        </template>
        <a-form :model="form" ref="formRef">
          <a-form-item field="key" label="键名" :rules="[{ required: true, message: '��输入键名' }]">
            <a-input v-model="form.key" placeholder="请输入键名" />
          </a-form-item>
          <a-form-item field="value" label="键值" :rules="[{ required: true, message: '请输入键值' }]">
            <a-input v-model="form.value" placeholder="请输入键值" />
          </a-form-item>
          <a-form-item field="type" label="类型" :rules="[{ required: true, message: '请选择类型' }]">
            <a-select v-model="form.type" placeholder="请选择类型">
              <a-option v-for="option in typeOptions" :key="option.value" :value="option.value">
                {{ option.label }}
              </a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="state" label="状态">
            <a-switch v-model="form.state" />
          </a-form-item>
          <a-form-item field="comment" label="名称">
            <a-textarea 
              v-model="form.comment" 
              placeholder="请输入名称" 
              :auto-size="{ minRows: 3, maxRows: 5 }"
            />
          </a-form-item>
        </a-form>
      </a-modal>
    </div>
</template>

<style>
.config-container {
  padding: 20px;
  margin-top: 60px;
  height: calc(100vh - 60px);
  overflow-y: auto;
  background-color: var(--color-bg-1);
}
#xgPlayerWrap { flex: auto; }
#xgPlayerWrap video { width: 100%; }
</style>

<script setup>
import { ref } from 'vue';
import { Message, Tag } from '@arco-design/web-vue';
import axios from 'axios';
import { onMounted } from 'vue';

const columns = [
  {
    title: '键名',
    dataIndex: 'key',
  },
  {
    title: '键值',
    dataIndex: 'value',
  },
  {
    title: '类型',
    dataIndex: 'type',
  },
  {
    title: '状态',
    dataIndex: 'state',
    slotName: 'state'
  },
  {
    dataIndex: 'comment',
    title: '名称',
  },
  {
    title: '操作',
    slotName: 'optional',
  },
];

const data = ref([]);
const visible = ref(false);
const editingRecord = ref(null);
const formRef = ref(null);
const typeOptions = ref([]);

const form = ref({
  key: '',
  value: '',
  type: '',
  state: true,
  comment: ''
});

// 获取字典类型选项
const fetchTypeOptions = () => {
  axios.post('/sys/config', {
    mt: 'queryList',
    type: 'DICT_TYPE'
  }).then(res => {
    if (res.data && res.code === 200) {
      typeOptions.value = res.data.map(item => ({
        label: item.key,
        value: item.value
      }));
      if (typeOptions.value.length === 0) {
        typeOptions.value = [{
          label: '字典类型',
          value: 'DICT_TYPE'
        }];
      }
    } else {
      typeOptions.value = [{
        label: '字典类型',
        value: 'DICT_TYPE'
      }];
      Message.error('获取字典类型失败');
    }
  }).catch(error => {
    console.error('获取字典类型失败:', error);
    typeOptions.value = [{
      label: '字典类型',
      value: 'DICT_TYPE'
    }];
    Message.error('获取字典类型失败');
  });
};

// 获取字典列表
const fetchDictList = () => {
  axios.post('/sys/config', {
    mt: 'queryList'
  }).then(res => {
    if (res.data && res.code === 200) {
      data.value = res.data;
    } else {
      Message.error('获取配置列表失败');
    }
  }).catch(error => {
    console.error('获取配置列表失败:', error);
    Message.error('获取配置列表失败');
  });
};

// 新增��钮点击
const handleAdd = () => {
  editingRecord.value = null;
  form.value = {
    key: '',
    value: '',
    type: 'DICT_TYPE',
    state: true,
    comment: ''
  };
  visible.value = true;
};

// 编辑按钮点击
const handleEdit = (record) => {
  editingRecord.value = record;
  form.value = {
    key: record.key,
    value: record.value,
    type: record.type,
    state: record.state,
    comment: record.comment
  };
  visible.value = true;
};

// 删除按钮点击
const handleDelete = (record) => {
  axios.post('/sys/config', {
    mt: 'delete',
    id: record.id
  }).then(res => {
    if (res.data && res.data.code === 200) {
      Message.success('删除成功');
      fetchDictList();
    } else {
      Message.error(res.data?.msg || '删除失败');
    }
  }).catch(error => {
    console.error('删除失败:', error);
    Message.error('删除失败');
  });
};

// 确认按钮点击
const handleOk = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      const params = {
        mt: 'add',
        ...form.value
      };
      if (editingRecord.value) {
        params.id = editingRecord.value.id;
      }
      
      axios.post('/sys/config', params).then(res => {
        if (res.data && res.data.code === 200) {
          Message.success(editingRecord.value ? '更新成功' : '添加成功');
          visible.value = false;
          fetchDictList();
        } else {
          Message.error(res.data?.msg || (editingRecord.value ? '更新失败' : '添加失败'));
        }
      }).catch(error => {
        console.error(editingRecord.value ? '更新失败' : '添加失败', error);
        Message.error(editingRecord.value ? '更新失败' : '添加失败');
      });
    }
  });
};

// 取消按钮点击
const handleCancel = () => {
  visible.value = false;
};

// 页面加载时初始化数据
onMounted(() => {
  fetchTypeOptions();
  fetchDictList();
});
</script>