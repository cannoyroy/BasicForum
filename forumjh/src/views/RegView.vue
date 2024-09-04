<template>
   <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />
    
    <div class="wrapper">
      <h1>论坛前端DEMO</h1>

      <nav>
        <RouterLink to="/login">登录</RouterLink>
        <RouterLink to="/reg">注册</RouterLink>
      </nav>
    </div>
  </header>
  <div class="register-container">
    <h1>注册界面</h1>
    <form @submit.prevent="register">
      <div class="form-group">
        <label for="username">用户名:</label>
        <input type="text" id="username" v-model="username" @input="limitCharacters_1" required>
      </div>
      <div class="form-group">
        <label for="name">姓名:</label>
        <input type="text" id="name" v-model="name" @input="limitCharacters_2" required>
      </div>
      <div class="form-group">
        <label for="password">密码:</label>
        <input type="password" id="password" v-model="password" @input="limitCharacters_3" placeholder="请输入8~16位密码" required>
      </div>
      <div class="form-group">
        <label for="user_type">用户类型:</label>
        <select id="user_type" v-model="user_type">
          <option value=1>普通用户</option>
          <option value=2>管理员</option>
        </select>
      </div>
      <button type="submit">注册</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      name: '',
      password: '',
      user_type: ''
    };
  },
  methods: {
    limitCharacters_1(event) {
      if (this.username.length > 50) {
        // 如果输入长度超过50，则截断输入值
        event.target.value = this.username.substring(0, 50);
        // 同步v-model绑定的数据
        this.username = event.target.value;
      }
    },
    limitCharacters_2(event) {
      if (this.name.length > 100) {
        event.target.value = this.name.substring(0, 100);
        // 同步v-model绑定的数据
        this.name = event.target.value;
      }
    },
    limitCharacters_3(event) {
      if (this.password.length > 16) {
        event.target.value = this.password.substring(0, 16);
        // 同步v-model绑定的数据
        this.password = event.target.value;
      }
    },
    register() {
      const userData = {
        username: this.username,
        name: this.name,
        password: this.password,
        user_type: parseInt(this.user_type)
      };

      axios.post("http://127.0.0.1:8080/api/user/reg", userData)
        .then(response => {
          // 处理注册成功的响应
          console.log(response.data);
          // 这里可以根据实际情况进行页面跳转或状态更新
          if (response.data.msg === "success") {
              // 弹出注册成功的消息
            alert("注册成功！");

              // 显示用户的相关信息
            // const userInfo = `用户名: ${response.data.username}\n姓名:${response.data.name}\n用户类型: ${response.data.user_type}`;
            // alert(userInfo);

              // 这里可以根据实际情况进行页面跳转或状态更新
          } else {
              // 如果msg不等于"success"，处理错误情况
            alert("注册失败\n"+response.data.msg);
          }
        })
        .catch(error => {
          // 处理注册失败的响应
          console.error(error);
        });
    }
  }
};
</script>

<style scoped>
@import './src/style/header.css';
</style>


<style>
.register-container {
  width: 80%;
  background-image: linear-gradient(to bottom, #c9e4ff, #ffffff);
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

.form-group select {
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
