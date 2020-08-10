<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增学生</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="姓名" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="性别" prop="sex" width="120">
         <template slot-scope="scope">{{scope.row.sex|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="年龄" prop="age" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateStudent(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteStudent(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

		<el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
			<el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
				<el-form-item label="姓名" prop="name">
					<el-input v-model="formData.name" placeholder="请输入姓名" clearable :style="{width: '100%'}"></el-input>
				</el-form-item>
				<el-form-item label="年龄" prop="age">
					<el-input-number v-model="formData.age"></el-input-number>
				</el-form-item>
				<el-form-item label="性别" prop="sex">
					<el-radio-group v-model="formData.sex" size="medium">
						<el-radio v-for="(item, index) in sexOptions" :key="index" :label="item.value"
											:disabled="item.disabled">{{item.label}}</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item size="large">
					<el-button type="primary" @click="submitForm">提交</el-button>
					<el-button @click="resetForm">重置</el-button>
				</el-form-item>
			</el-form>

			<div class="dialog-footer" slot="footer">
				<el-button @click="closeDialog">取 消</el-button>
				<el-button @click="enterDialog" type="primary">确 定</el-button>
			</div>
		</el-dialog>
  </div>
</template>

<script>
import {
    createStudent,
    deleteStudent,
    deleteStudentByIds,
    updateStudent,
    findStudent,
    getStudentList
} from "@/api/student";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Student",
  mixins: [infoList],
  data() {
    return {
      listApi: getStudentList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
			formData: {
				name: undefined,
				age: undefined,
				sex: 0,
			},
			rules: {
				name: [{
					required: true,
					message: '请输入姓名',
					trigger: 'blur'
				}],
				age: [{
					required: true,
					message: '',
					trigger: 'blur'
				}],
				sex: [{
					required: true,
					message: '性别不能为空',
					trigger: 'change'
				}],
			},
			sexOptions: [{
				"label": "男",
				"value": 0
			}, {
				"label": "女",
				"value": 1
			}],
		}
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10       
        if (this.searchInfo.sex==""){
          this.searchInfo.sex=null
        }       
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteStudentByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateStudent(row) {
      const res = await findStudent({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.restudent;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          name:null,
          sex:null,
          age:null,
      };
    },
    async deleteStudent(row) {
      this.visible = false;
      const res = await deleteStudent({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createStudent(this.formData);
          break;
        case "update":
          res = await updateStudent(this.formData);
          break;
        default:
          res = await createStudent(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>
