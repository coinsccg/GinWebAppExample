<template>
  <div style="width:600px;background-color: white;height:600px;">
    <div class="userInfos_div">
      <div
        v-for="userInfot in userInfos"
        :key="userInfot.user.userId"
        @click="userInfo=userInfot"
        class="userInfos_div_div"
      >
        <div>
          <userInfoSmall :userInfo="userInfot" style="margin-top:10px;"></userInfoSmall>
        </div>
      </div>
    </div>
    <el-card v-if="userInfo" :body-style="{ padding: '10px' }" class="chat_card">
      <div slot="header">
        <span>{{userInfo.user.nickname}}</span>
      </div>
      <div ref="messages" class="message_div">
        <div v-for="(message,index) in messages" :key="index">
          <div
            :style="messageStyle(message.dstId===userId)"
            style="text-align:left;margin-top:10px;"
          >
            <div>{{message.dstId===userId?userInfo.user.nickname:nickname}} {{$getTimeString(message.gmtCreate)}}</div>
            <div>{{message.msg}}</div>
          </div>
        </div>
      </div>
      <div style="margin-top:10px;">
        <el-input v-model="msg" placeholder="聊天内容" style="float:left;width:300px;"></el-input>
        <el-button @click="handleSendMessage" style="float:right;">发送</el-button>
      </div>
    </el-card>
  </div>
</template>
<script>
import userInfoSmall from "../components/UserInfoSmall";
import { queryUserInfos } from "@/api/user";
export default {
  name: "Chat",
  data() {
    return {
      msg: "",
      userInfo: null,
      userInfos: [],
      messages: []
    };
  },
  components: {
    userInfoSmall
  },
  methods: {
    handleSendMessage() {
      const message = {
        dstId: this.dstId,
        msg: this.msg
      };
      this.chatWebSocket.send(JSON.stringify(message));
      this.msg = "";
      this.messages = this.$store.getters.messageGroup.get(this.dstId);
    },
    messageStyle(val) {
      return {
        color: val ? "blue" : "green"
      };
    },
    updateMessages() {
      const messageGroup = this.$store.getters.messageGroup;
      const peopleIds = [];
      for (const [userId, messages] of messageGroup) {
        peopleIds.push(userId);
      }
      queryUserInfos(peopleIds).then(res => {
        this.userInfos = res;
      });
      if (this.dstId)
        this.messages = this.$store.getters.messageGroup.get(this.dstId);
    }
  },
  created() {
    if (this.peopleUserInfo) {
      this.userInfo = this.peopleUserInfo;
    }
    this.updateMessages();
  },
  computed: {
    userId() {
      return this.$store.getters.userId;
    },
    chatWebSocket() {
      return this.$store.getters.chatWebSocket;
    },
    dstId() {
      return this.userInfo ? this.userInfo.user.userId : null;
    },
    nickname() {
      return this.$store.getters.user.nickname;
    },
    valueChange() {
      return this.$store.getters.valueChange;
    }
  },
  watch: {
    dstId(val) {
      if (val) this.messages = this.$store.getters.messageGroup.get(val);
    },
    valueChange(val) {
      this.updateMessages();
    },
    messages(val) {
      setTimeout(() => {
        var obj = this.$refs.messages;
        obj.scrollTop = obj.scrollHeight;
      }, 500);
    }
  },
  props: ["peopleUserInfo"]
};
</script>
<style scoped>
.chat_card {
  width: 398px;
  height: 598px;
  float: left;
}
.message_div {
  height: 470px;
  overflow: auto;
}
.userInfos_div {
  overflow: auto;
  float: left;
  background-color: white;
  height: 600px;
  width: 200px;
}
.userInfos_div_div {
  margin-left: 10px;
  height: 50px;
  clear: both;
  cursor: pointer;
}
</style>