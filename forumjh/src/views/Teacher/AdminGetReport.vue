<template>
  
    <div class="post-container">
      <div class="post-list-wrapper">
        <div class="post-list">
          <!-- 帖子列表将在这里动态渲染 -->
          <div v-for="post in posts" :key="post.id" class="post">
            <div class="user-info">
              帖子编号：{{ post.PostID }}  |  举报的用户名: {{post.Username}}
            </div>
            <div class="content">
                举报内容：
              {{ post.Content }}
            </div>
            <div class="content">
              举报原因：
              {{ post.Reason }}
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
    if (route.query.user_type === '2') {
      console.log('User type is 2, proceeding with page mount.');
    }else{
      alert("请确认访问权限");
      window.location.href = "http://localhost:5173/login";
    }
    try {
      const response = await axios.get('http://127.0.0.1:8080/api/admin/report', {
        params: { user_id: route.query.user_id } // 假设 user_id 来自查询参数
      });
  
      posts.value = response.data.data.report_list;
      console.log(response.data.data)
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
};</script>
  
  
  <style scoped>
  @import './src/style/header.css';
  </style>
  
  
<style scoped>
.post-container {
  max-width: 30vw; /* 设置最大宽度为屏幕宽度的50% */
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
  overflow-wrap: break-word; /* 允许单词在到达边界时断开 */
  word-wrap: break-word; /* 允许单词在到达边界时断开 */
}
</style>