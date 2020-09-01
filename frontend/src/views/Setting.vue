<template>
  <div>
    <el-card>
      <el-upload
        action="/api/uploadAvater"
        :headers="headers"
        :show-file-list="false"
        :on-success="uploadAvaterSuccess"
        accept=".png, .jpg"
      >
        <img :src="avaterUrl" class="user_avater" />
      </el-upload>
      <el-divider>
        <span>INFOMATION</span>
      </el-divider>
      <el-form
        :rules="userRules"
        label-position="left"
        :model="userForm"
        ref="userForm"
        label-width="100px"
        status-icon
      >
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="userForm.nickname"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userForm.email"></el-input>
        </el-form-item>
        <el-form-item label="性别" prop="sex">
          <el-select v-model="userForm.sex">
            <el-option label="男" :value="0"></el-option>
            <el-option label="女" :value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="简介" prop="intro">
          <el-input type="textarea" v-model="userForm.intro"></el-input>
        </el-form-item>
        <el-button type="primary" @click="handleUpdateUser">保存</el-button>
        <el-button @click="resetUser">重置</el-button>
      </el-form>
      <el-divider>
        <span style="color:red">DANGER</span>
      </el-divider>
      <el-form
        :rules="passwordRules"
        label-position="left"
        :model="passwordForm"
        ref="passwordForm"
        label-width="100px"
        status-icon
      >
        <el-form-item label="旧密码" prop="oldPassword">
          <el-input type="password" v-model="passwordForm.oldPassword"></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input type="password" v-model="passwordForm.newPassword"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="passwordConfirm">
          <el-input type="password" v-model="passwordForm.passwordConfirm"></el-input>
        </el-form-item>
        <el-button type="primary" @click="handleUpdatePassword">修改密码</el-button>
      </el-form>
    </el-card>
  </div>
</template>
<script>
import { updateUser, updatePassword, uploadAvater } from "@/api/user";
export default {
  name: "Setting",
  data() {
    return {
      userForm: {
        nickname: "",
        sex: 0,
        intro: "",
        email: ""
      },
      passwordForm: {
        oldPassword: "",
        newPassword: "",
        passwordConfirm: ""
      },
      userRules: {},
      passwordRules: {
        oldPassword: [
          {
            required: true,
            message: "密码不能为空"
          }
        ],
        newPassword: [
          {
            required: true,
            message: "密码不能为空"
          },
          { min: 3, max: 15, message: "长度在3到15个字符" }
        ],
        passwordConfirm: [
          {
            required: true,
            message: "请确认密码"
          },
          {
            validator: (rule, value, callback) => {
              if (value !== this.passwordForm.newPassword)
                callback(new Error("密码输入不一致"));
              callback();
            },
            trigger: "change"
          }
        ]
      }
    };
  },
  created() {
    this.resetUser();
  },
  methods: {
    handleUpdateUser() {
      this.$refs.userForm.validate(vaild => {
        if (vaild) {
          const data = JSON.parse(JSON.stringify(this.userForm));
          updateUser(data).then(res => {
            if (res === true) {
              this.$message({
                showClose: true,
                message: "资料修改成功",
                type: "success"
              });
            }
            this.$store.commit("setInfo", data);
          });
        }
      });
    },
    handleUpdatePassword() {
      this.$refs.passwordForm.validate(vaild => {
        if (vaild) {
          const data = {
            oldPassword: this.passwordForm.oldPassword,
            newPassword: this.passwordForm.newPassword
          };
          updatePassword(data).then(res => {
            if (res === true) {
              this.$message({
                showClose: true,
                message: "密码修改成功",
                type: "success"
              });
            }
          });
        }
      });
    },
    resetUser() {
      const user = this.$store.getters.user;
      this.userForm.nickname = user.nickname;
      this.userForm.sex = user.sex;
      this.userForm.intro = user.intro;
      this.userForm.email = user.email;
    },
    uploadAvaterSuccess(val) {
      this.$store.commit("setAvater");
    }
  },
  computed: {
    user() {
      return this.$store.getters.user;
    },
    avaterUrl() {
      return this.$avaterUrl(this.user.avater, this.user.userId);
    },
    headers() {
      return {
        token: this.$store.getters.token
      };
    }
  }
};
</script>
<style scoped>
.user_avater {
  width: 160px;
  padding: 4px;
  height: 160px;
  border-radius: 5%;
}
</style>