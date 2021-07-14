<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>Sign in</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Sign in</ion-title>
        </ion-toolbar>
      </ion-header>
    
      
          <ion-button @click="onOAuthBtnClick()">Sign in</ion-button>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { OAuth2Client } from "@amischencko/capacitor-oauth2";
import { IonPage, IonHeader, IonToolbar, IonButton } from '@ionic/vue';

export default  {
  components: { IonHeader, IonToolbar, IonPage, IonButton },
  methods: { 
    onOAuthBtnClick() {
      console.log("click");
        OAuth2Client.authenticate( 
          {
          "scope": "email profile", 
            "web":{
              "appId":"974961430260-rikreoe46uh27vggl9989ksq6r7ird8u.apps.googleusercontent.com", 
              //"accessTokenEndpoint": "https://oauth2.googleapis.com/token", 
              "authorizationBaseUrl": "https://accounts.google.com/o/oauth2/auth", 
              "responseType": "token", 
              "redirectUrl": "http://localhost:8100/tabs/tab1"
            }
          }
        ).then(response => {
            const accessToken = response["access_token"];
            const refreshToken = response["refresh_token"];
            const oauthUserId = response["id"];
            const name = response["name"];

            localStorage.setItem("accessToken", accessToken);

            // go to backend
        }).catch(reason => {
            console.error("OAuth rejected", reason);
        });
    }
  }
}
</script>