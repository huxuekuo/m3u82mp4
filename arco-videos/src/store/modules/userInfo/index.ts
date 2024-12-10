import { defineStore } from 'pinia';

const useUserInfoStore = defineStore('ui', {
    state:()=>({
        avatar : ""
    }),
    actions:{
        SetAvatar(avatar:string){
            this.avatar = avatar
        }
    }
})

export default useUserInfoStore;