<template>
  <div>
    <el-form
      :model="questionForm"
      :rules="questionRules"
      class="questionForm"
      label-width="100px"
      label-position="left"
      ref="questionForm"
    >
      <el-form-item label="标题" prop="caption">
        <el-input v-model="questionForm.caption"></el-input>
      </el-form-item>
      <el-form-item label="问题描述" prop="content">
        <el-input type="textarea" v-model="questionForm.content"></el-input>
      </el-form-item>
      <el-form-item label="分类" prop="categoryList">
        <el-select v-model="questionForm.category_id" placeholder="请选择" style="display:block">
          <el-option
            v-for="category in categoryList"
            :key="category.id"
            :label="category.name"
            :value="category.id"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-button v-if="hasQuestion" type="primary" @click="handleUpdateQuestion"
        >修改问题</el-button
      >
      <el-button v-else type="primary" @click="handleSaveQuestion"
        >提问</el-button
      >
    </el-form>
  </div>
</template>
<script>
import { saveQuestion, updateQuestion } from "@/api/question";
import { getCategoryList } from "@/api/category";
export default {
  name: "SaveQuestion",
  data() {
    return {
      categoryList: [],
      inputVisible: false,
      questionForm: {
        caption: "",
        content: "",
        category_id: null
      },
      questionRules: {
        caption: [
          {
            required: true,
            message: "问题标题不能为空"
          }
        ]
      }
    };
  },
  created() {
    if (this.hasQuestion) {
      const question = this.question;
      const questionForm = this.questionForm;
      questionForm.caption = question.caption;
      questionForm.content = question.content;
    }
    getCategoryList().then(res => {
      console.log(res);
      if (res.code !== 1000) {
        this.$message({
          showClose: true,
          message: res.message,
          type: "error"
        });
      } else {
        this.categoryList = res.data;
      }
    });
  },
  methods: {
    handleSaveQuestion() {
      this.$refs.questionForm.validate(vaild => {
        if (vaild) {
          const data = JSON.parse(JSON.stringify(this.questionForm));
          saveQuestion(data).then(res => {
            console.log(res.data)
            if (res.data > 0) {
              this.$router.push({
                path: "/question",
                query: {
                  question_id: res.data
                }
              });
            }
          });
        }
      });
    },
    handleUpdateQuestion() {
      this.$refs.questionForm.validate(vaild => {
        if (vaild) {
          const data = JSON.parse(JSON.stringify(this.questionForm));
          data["questionId"] = this.question.questionId;
          updateQuestion(data).then(res => {
            if (res === true) {
              this.$emit("updateQuestion", data);
            }
          });
        }
      });
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.topic.$refs.input.focus();
      });
    }
  },
  computed: {
    hasQuestion() {
      return this.question != null;
    }
  },
  props: ["question"]
};
</script>
<style scoped>
.questionForm {
  text-align: center;
}
.topic_input {
  width: 70px;
}
</style>
