<script setup>
import { useRoute, RouterLink, RouterView } from 'vue-router';
const route = useRoute();
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
      </nav>
    </div>
  </header>

  <div class="post-container">
    <form @submit.prevent="post">
        <div class="form-group">
        <!-- <label for="content">内容:</label> -->
        <textarea id="post_id" v-model="post_id" required placeholder= "请输入要举报的帖子ID" style="height: 50px;"></textarea>
        <textarea id="content" v-model="content" required placeholder="请输入举报原因"></textarea>
      </div>
      <button type="submit">举报</button>
    </form>
  </div>

  <RouterView />
</template>

<script>
import axios from 'axios';

export default {
  computed: {
    userInfo() {
      return this.$route.query;
    }
  },
  data() {
    return {
      content: '',
      post_id: ''
    };
  },
  methods: {
    post() {
      const query = this.$route.query;
      const postData = {
        reason: this.content,
        user_id: parseInt(query.user_id), // 这里应该从全局状态或登录状态中获取用户 ID
        post_id: parseInt(this.post_id)
      };

    //   console.log(this.content)

      axios.post("http://127.0.0.1:8080/api/student/report-post", postData)
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("举报成功")
          this.content = ''
          this.post_id = ''
        })
        .catch(error => {
          console.error('Fail', error);
          alert("无法举报，请稍后再试。")
        });
    },

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
  width: 80%;
  margin: auto;
  padding: 20px;
  background-color: #f0f0f0;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 15px;
}

.form-group textarea {
  width: 100%;
  height: 200px; /* 固定输入框大小 */
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  resize: none; /* 禁止拖拽缩放 */
}

button {
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 4px;
  background-color: #ff9191; /* 自定义背景颜色 */
  color: rgb(56, 56, 56); /* 自定义文字颜色 */
  cursor: pointer;
}

button:hover {
  background-color: #ffd8d8; /* 鼠标悬停时的背景颜色 */
}
</style>