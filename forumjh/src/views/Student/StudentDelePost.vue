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

  <div class="post-container">
    <form @submit.prevent="post">
      <div class="form-group">
        <textarea id="post_id" v-model="post_id" required placeholder= "请输入帖子ID"></textarea>
      </div>
      <button type="submit">删除</button>
    </form>
  </div>

  <RouterView />
</template>


<script>
import axios from 'axios';

export default {
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
        post_id: parseInt(this.post_id),
        user_id: parseInt(query.user_id) // 这里应该从全局状态或登录状态中获取用户 ID
      };

    //   console.log(postData)

      axios.delete("http://127.0.0.1:8080/api/student/post", {params:postData})
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("帖子删除成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          this.post_id = ''
        })
        .catch(error => {
          console.error('Fail', error);
          alert("无法删除，请确认您想删除的帖子是属于您的或者帖子是否存在。")
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
  height: 50px; /* 固定输入框大小 */
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
  color: rgb(10, 8, 8); /* 自定义文字颜色 */
  cursor: pointer;
}

button:hover {
  background-color: #ffd8d8; /* 鼠标悬停时的背景颜色 */
}
</style>