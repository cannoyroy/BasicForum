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
        <!-- <label for="content">内容:</label> -->
        <textarea id="content" v-model="content"  @input="updateCounter" required placeholder="请输入修改后的内容，不得超过255字"></textarea>
        <span class="counter" :style="counterStyle">字数：{{ characterCount }}/ 255</span>
        <textarea id="post_id" v-model="post_id" required placeholder= "请输入帖子ID" style="height: 50px;"></textarea>
      </div>
      <button type="submit">修改</button>
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
      post_id: '',
      characterCount: 0,
    };
  },
  computed: {
    counterStyle() {
      return {
        color: this.characterCount > 255 ? 'red' : 'black',
      };
    },
  },
  methods: {
    updateCounter() {
      this.characterCount = this.content.length;
    },
    post() {
      const query = this.$route.query;
      const postData = {
        content: this.content,
        post_id: parseInt(this.post_id),
        user_id: parseInt(query.user_id) // 这里应该从全局状态或登录状态中获取用户 ID
      };

      console.log(postData)

      axios.put("http://127.0.0.1:8080/api/student/post", postData)
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("帖子修改成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          this.content = ''
          this.post_id = ''
        })
        .catch(error => {
          console.error('Fail', error);
          alert("无法修改，请确认您想修改的帖子是属于您的或者帖子是否存在。")
        });
    },

    goToSpecifiedLink() {
      window.location.href = "http://localhost:5173/login"; // 这里替换成你要跳转的链接
    }
  },
  watch: {
    inputText(newVal) {
      this.updateCounter();
    },
  },
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
  background-color: #9ee9ae; /* 自定义背景颜色 */
  color: rgb(56, 56, 56); /* 自定义文字颜色 */
  cursor: pointer;
}

button:hover {
  background-color: #daffcf; /* 鼠标悬停时的背景颜色 */
}
</style>