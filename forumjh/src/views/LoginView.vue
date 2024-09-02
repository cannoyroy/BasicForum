<template>
   <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" @click="goToSpecifiedLink">
    
    <div class="wrapper">
      <h1>论坛前端DEMO</h1>

      <nav>
        <RouterLink to="/login">登录</RouterLink>
        <RouterLink to="/reg">注册</RouterLink>
      </nav>
    </div>
  </header>
  <div class="login-container">
    <h1>登录界面</h1>
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="username">学号/管理工号:</label>
        <input type="text" id="username" v-model="username" required>
      </div>
      <div class="form-group">
        <label for="password">密码:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <button type="submit">登录</button>
    </form>
  </div>

</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    login() {
      const userData = {
        username: this.username,
        password: this.password
      };


      axios.post("http://127.0.0.1:8080/api/user/login", userData)
        .then(response => {
          // 处理登录成功的响应
          console.log(response.data);
          // 这里可以根据实际情况进行页面跳转或状态更新
          if (response.data.msg === "success") {

            const userInfo = {
              user_id: response.data.data.user_id,
              username: response.data.data.username,
              name: response.data.data.name,
              user_type: response.data.data.user_type,
            };
            console.log(userInfo)
            if (response.data.data.user_type == 1) {
              alert("登陆成功\n点击图标随时返回主界面")
              this.$router.push({ name: 'StudentBoard', query: userInfo});
            }else{
              alert("登陆成功\n点击图标随时返回主界面")
              this.$router.push({ name: 'TeacherBoard', query: userInfo});
            }

          } else {
            alert("登录失败\n"+response.data.msg);
          }
        })
        .catch(error => {
          // 处理登录失败的响应
          console.error(error);
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


<style>
.login-container {
  width: 80%;
  background-image: linear-gradient(to bottom, #ffdfdf, #ffffff);
  max-width: 300px;
  min-width: 50px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-group input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
