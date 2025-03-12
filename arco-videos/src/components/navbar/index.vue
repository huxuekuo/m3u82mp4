<template>
  <div class="navbar">
    <div class="left-side">
      <a href="/query">
      <a-space>
        <img
          alt="logo"
          src="@/assets/images/2.png"
        />
        <a-typography-title
          :style="{ margin: 0, fontSize: '18px' }"
          :heading="5"
        >
        追剧神器
        </a-typography-title>
      </a-space>
    </a>
    </div>
    <div class="center-side">
      <Menu v-if="topMenu" />
    </div>
    <ul class="right-side">
      <!-- <li>
        <a-tooltip
          content="系统配置"
        >
          <a-button
            class="nav-btn"
            type="outline"
            shape="circle"
            href="config"
          >
            <template #icon>
              <icon-font type="icon-xitong" size="15"/>
            </template>
          </a-button>
        </a-tooltip>
      </li> -->
      <li>
        <a-tooltip
          :content="
            theme === 'light'
              ? $t('settings.navbar.theme.toDark')
              : $t('settings.navbar.theme.toLight')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="handleToggleTheme"
          >
            <template #icon>
              <icon-moon-fill v-if="theme === 'dark'" />
              <icon-sun-fill v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <a-tooltip
          :content="
            isFullscreen
              ? $t('settings.navbar.screen.toExit')
              : $t('settings.navbar.screen.toFull')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="toggleFullScreen"
          >
            <template #icon>
              <icon-fullscreen-exit v-if="isFullscreen" />
              <icon-fullscreen v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <a-tooltip content="留言板">
          <router-link to="/message">
            <a-button
              class="nav-btn"
              type="outline"
              shape="circle"
            >
              <template #icon>
                <icon-message />
              </template>
            </a-button>
          </router-link>
        </a-tooltip>
      </li>
      <li>
        <a-dropdown trigger="click">
          <a-avatar
            :size="32"
            :style="{ marginRight: '8px', cursor: 'pointer' }"
          >
            <img alt="avatar" src="@/assets/images/default-avatar.png" v-if="avatar == ''" />
            <img alt="avatar" :src="avatar" v-if="avatar != ''" />
          </a-avatar>
          <template #content>
            <a-doption>
              <a-space @click="selfInfo">
                <icon-export />
                <span>
                  个人信息
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="handleLogout">
                <icon-export />
                <span>
                  {{ $t('messageBox.logout') }}
                </span>
              </a-space>
            </a-doption>
          </template>
        </a-dropdown>
      </li>
    </ul>
    
  </div>
  
</template>

<script lang="ts" setup>
  import { computed, ref, inject } from 'vue';
  import { useDark, useToggle, useFullscreen } from '@vueuse/core';
  import { useAppStore, useUserStore,useUserInfoStore } from '@/store';
  import { Icon } from '@arco-design/web-vue';
  import useUser from '@/hooks/user';
  import Menu from '@/components/menu/index.vue';
  import UserInfo from '../user/userinfo.vue'

  const appStore = useAppStore();
  const userInfoStore = useUserInfoStore()
  const visibleParent = ref(0)
  const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/c/font_4690348_1mxb9zlmcwj.js' });
  const { logout } = useUser();
  const { isFullscreen, toggle: toggleFullScreen } = useFullscreen();
  const avatar = computed(() => {
    return userInfoStore.avatar;
  });
  const theme = computed(() => {
    return appStore.theme;
  });

  const topMenu = computed(() => appStore.topMenu && appStore.menu);
  const isDark = useDark({
    selector: 'body',
    attribute: 'arco-theme',
    valueDark: 'dark',
    valueLight: 'light',
    storageKey: 'arco-theme',
    onChanged(dark: boolean) {
      // overridden default behavior
      appStore.toggleTheme(dark);
    },
  });
  const toggleTheme = useToggle(isDark);
  const handleToggleTheme = () => {
    toggleTheme();
  };
  const handleLogout = () => {
    logout();
  };

  // 个人信息
  const selfInfo = () => {
    
  };

  function userInfoSettings(){
    visibleParent.value+=1
  }
  

</script>

<style scoped lang="less">
  .navbar {
    display: flex;
    justify-content: space-between;
    height: 100%;
    background-color: var(--color-bg-2);
    border-bottom: 1px solid var(--color-border);
  }

  .left-side {
    display: flex;
    align-items: center;
    padding-left: 20px;
  }

  .center-side {
    flex: 1;
  }

  .right-side {
    display: flex;
    padding-right: 20px;
    list-style: none;
    :deep(.locale-select) {
      border-radius: 20px;
    }
    li {
      display: flex;
      align-items: center;
      padding: 0 10px;
    }

    a {
      color: var(--color-text-1);
      text-decoration: none;
    }
    .nav-btn {
      border-color: rgb(var(--gray-2));
      color: rgb(var(--gray-8));
      font-size: 16px;
    }
    .trigger-btn,
    .ref-btn {
      position: absolute;
      bottom: 14px;
    }
    .trigger-btn {
      margin-left: 14px;
    }
  }
</style>

<style lang="less">
  .message-popover {
    .arco-popover-content {
      margin-top: 0;
    }
  }
</style>
