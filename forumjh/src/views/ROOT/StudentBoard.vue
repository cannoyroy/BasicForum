<script setup>
import { useRoute, RouterLink, RouterView } from 'vue-router';
import {onMounted} from 'vue';

const route = useRoute();


onMounted(async () => {
  if (route.query.user_type === '1') {
    console.log('User type is 1, proceeding with page mount.');
  }else{
    alert("请确认访问权限");
    window.location.href = "http://localhost:5173/login";
  }
});
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" @click="goToSpecifiedLink">
    
    <div class="wrapper">
      <h1>学生端界面</h1>

      <nav>
        <RouterLink :to="{ name: 'studentpost', query : route.query}">发帖</RouterLink>
        <RouterLink :to="{ name: 'studentgetpost', query : route.query}">查看</RouterLink>
        <RouterLink :to="{ name: 'studentputpost', query : route.query}">修改</RouterLink>
        <RouterLink :to="{ name: 'studentdelepost', query : route.query}">删除</RouterLink>
        <RouterLink :to="{ name: 'studentreportpost', query : route.query}">举报</RouterLink>
        <RouterLink :to="{ name: 'studentgetreport', query : route.query}">查看举报</RouterLink>
        <RouterLink :to="{ name: 'logon', query : {}}" style="color: red;">注销</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<script>
export default {
  methods: {
    goToSpecifiedLink() {
      window.location.href = "http://localhost:5173/login"; // 这里替换成你要跳转的链接
    }
  }
};
</script>

<style scoped>
@import './src/style/header.css';
</style>
