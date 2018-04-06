<template>
<div id="content">
  <div id="title">TODOLIST</div>
  <input type="text" hint="输入你计划做的事" id="input_text" v-model="newItem" @keyup.enter="onEnter">
  <div id="list">
    <div id="todolist">
      <div class="list_name">TODO</div>
      <ul type="square">
        <li v-for="item in items" :key="item" v-if="!item.IsFinished">
          <div class="label">{{item.Label}}</div>
          <input type="button" value="Finished" class="state" @click="doThing(item)">
        </li>
      </ul>
    </div>
    <div id="finishedlist">
      <div class="list_name">DOWN</div>
      <ul type="square">
        <li v-for="downitem in items" :key="downitem" v-if="downitem.IsFinished">
          <div class="label">{{downitem.Label}}</div>
          <input type="button" value="Delete" class="state" @click="doDelete(downitem)">
        </li>
      </ul>
    </div>
  </div>
</div>
</template>

<script>
import Resource from 'vue-resource'
import Vue from 'vue'
Vue.use(Resource)

export default {
  data: function () {
    return {
      items: [],
      newItem: ''
    }
  },
  created: function () {
    this.$http.get('/data', 'GET')
      .then((response) => {
        this.items = response.data
      })
      .catch(function (err) {
        console.log(err)
      })
  },
  methods: {
    pushData: function (item, vueIns) {
      vueIns.$http.post('/data', item, 'POST')
        .then((response) => {
          console.log(response.data)
        })
    },
    doThing: function (item) {
      item.IsFinished = !item.IsFinished
      item.method = 'update'
      this.$options.methods.pushData(item, this)
    },
    doDelete: function (downitem) {
      var index = this.items.find(function (value) {
        return value.Label === downitem.Label
      })
      this.items.splice(index, 1)
      downitem.method = 'delete'
      this.$options.methods.pushData(downitem, this)
    },
    onEnter: function () {
      var date = new Date()
      var item = {
        Label: this.newItem,
        IsFinished: false,
        CreateAt: date
      }
      this.items.push(item)
      item['method'] = 'add'
      this.$options.methods.pushData(item, this)
      this.newItem = ''
    }
  }
}
</script>

<style>
#content {
  margin: 100px auto 0 auto;
  width: 70%;
  height: calc(100% - 100px);
}
#title {
  font-size: 100px;
  text-align: center;
  color: #4285F8;
}
#input_text {
  height: 40px;
  width: 90%;
  margin: 30px 5% 0 5%;
}
#list {
  margin: 70px auto;
  height: calc(100% - 345px);
  width: 100%;
  position: relative;
}
#todolist, #finishedlist {
  padding:0;
  width: 40%;
  height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
  border: #4285F8 2px solid;
}
#todolist {
  position: absolute;
  margin-left: 8%;
  margin-right: 2%;
}
#finishedlist {
  position: absolute;
  margin-left: 52%;
}
.list_name {
  margin: 20px auto 0 auto;
  font-size: 30px;
  font-weight: bolder;
  color: #E6A98A;
}
ul, li {
  position: relative;
  margin: auto;
  text-align: left;
}
ul {
  list-style-position: outside;
  padding:10px 20px 10px 40px;
  width: calc(100%-60px);
  height: calc(100%-20px);
}
li {
  width: 100%;
  height: 40px;
  color:gray;
  font-weight: bolder;
}
.label {
  height: 30px;
  top: 0;
  left: 0;
  position: absolute;
  font-size: 20px;
  margin: 0;
  width:60%;
  height: 30px;
}
.state {
  position: absolute;
  margin-left: 70%;
  height: 30px;
  width:95px;
  background-color: rgb(240, 149, 104);
  color:#b45424;
  font-weight: bolder;
  border: none;
  border-radius: 5px;
}
</style>
