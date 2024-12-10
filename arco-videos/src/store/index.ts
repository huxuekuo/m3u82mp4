import { createPinia } from 'pinia';
import useAppStore from './modules/app';
import useUserStore from './modules/user';
import useTabBarStore from './modules/tab-bar';
import useUserInfoStore from './modules/userInfo'

const pinia = createPinia();

export { useAppStore, useUserStore, useTabBarStore,useUserInfoStore};
export default pinia;
