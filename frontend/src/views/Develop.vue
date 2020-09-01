<template>
  <div>
    <el-tooltip :content="flushStr" placement="bottom" effect="light">
      <el-button @click="handleFlush" plain>清空</el-button>
    </el-tooltip>
    <el-tooltip :content="backupStr" placement="bottom" effect="light">
      <el-button @click="handleBackup" plain>备份</el-button>
    </el-tooltip>
    <el-tooltip :content="restoreStr" placement="bottom" effect="light">
      <el-button @click="handleRestore" plain>恢复</el-button>
    </el-tooltip>
    <el-tooltip content="search" placement="bottom" effect="light">
      <el-button @click="handleCreateIndex" plain>重新建立索引</el-button>
    </el-tooltip>
  </div>
</template>
<script>
import { restore, flush, backup, createIndex } from "@/api/develop";
export default {
  data() {
    return {
      flushStr:
        "清空redis中的所有数据。所有关系将会丢失，索引也将丢失。请将redis数据被备份到mysql后再使用。",
      backupStr:
        "将mysql中的活动备份到各个表中。活动包括用户对回答的赞同、关注用户、关注问题、关注话题。",
      restoreStr:
        "从mysql中恢复问题、回答、评论之间id关系。将mysql中的活动数据恢复到redis。活动包括用户对回答的赞同、关注用户、关注问题、关注话题。"
    };
  },
  methods: {
    handleCreateIndex() {
      createIndex().then(res => {
        console.log(res);
      });
    },
    handleRestore() {
      restore().then(res => {
        console.log(res);
      });
    },
    handleFlush() {
      flush().then(res => {
        console.log(res);
      });
    },
    handleBackup() {
      backup().then(res => {
        console.log(res);
      });
    }
  },
  created() {
    console.log([1, 2, 3].indexOf(4));
  }
};
</script>