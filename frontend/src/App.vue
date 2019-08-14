<template>
  <div v-if="user">
  	<button v-on:click="login">Please login</button>;
  </div>
  <div v-else>
    <div>{{message}}</div>
    <p style="color:red;">{{errorMessage}}</p>
    <button v-on:click="getPrivateMessage">
      Get Private Message
    </button>
    <button v-on:click="logout">Logout</button>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import firebase from "./firebase";
import { getPrivateMessage, getPublicMessage } from "./api";

export default Vue.extend({

  created() {
  	firebase.auth().onAuthStateChanged(user => {
  	  if (user) {
  	    this.user = user ;
  	  } else {
  	    this.user = null
  	  }
  	});
  },
  data() {
    return {
      user: null,
      message: "",
      errorMessage: ""
    }
  },
  methods: {

  	login: function() {
  		firebase.login();
  	},

  	logout: function() {
  		firebase.logout();
  	},

  	getPrivateMessage: function() {
  		this.user
  		  .getIdToken()
  		  .then(token => {
  		    return getPrivateMessage(token);
  		  })
  		  .then(resp => {
  		    this.message = resp.message;
  		  })
  		  .catch(error => {
  		    this.errorMessage = error.toString()
  		  });
  	}
  }
});
</script>

<style lang="scss" scoped>
.container {
  color: green;
}
</style>
