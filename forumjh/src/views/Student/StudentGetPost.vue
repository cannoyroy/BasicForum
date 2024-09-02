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
      </nav>
    </div>
  </header>

  <div class="post-container">
    <div class="post-list-wrapper">
      <div class="post-list">
        <!-- 帖子列表将在这里动态渲染 -->
        <div v-for="post in posts" :key="post.id" class="post">
          <div class="user-info">
            {{ post.PostID }}  |  user: {{ post.UserID }}  |  发布时间: {{ post.CreatedAt.slice(0, 10)+" "+post.CreatedAt.slice(11, 16) }}
          </div>
          <div class="content">
            {{ post.Content }}
          </div>
        </div>
      </div>
    </div>
  </div>

  <RouterView />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, RouterLink, RouterView } from 'vue-router';

const route = useRoute();
const posts = ref([]);

onMounted(async () => {
  try {
    const response = await axios.get('http://127.0.0.1:8080/api/student/post', {
      params: { user_id: route.query.user_id } // 假设 user_id 来自查询参数
    });

    posts.value = response.data.data.post_list;
    // console.log(response.data.data.post_list)
  } catch (error) {
    console.error('Failed to fetch posts:', error);
  }
});


</script>

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


<style scoped>
.post-container {
  /* 样式代码 */
}

.post-list-wrapper {
  max-height: 400px; /* 限定框框的高度 */
  overflow-y: auto; /* 添加滚轮滑动功能 */
  border: 1px solid #ccc;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.post-list {
  padding: 10px;
}

.post {
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
}

.user-info {
  font-weight: bold;
  margin-bottom: 5px;
}

.content {
  margin-bottom: 10px;
}
</style>