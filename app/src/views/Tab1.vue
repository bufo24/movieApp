<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>Tab 1</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Swipe</ion-title>
        </ion-toolbar>
      </ion-header>
    
      <ExploreContainer name="Tab 1 page" />

      <ion-body>
        <div class="redline"></div>
      </ion-body>

      <ion-grid>
        <ion-item>
          <ion-row>
            <ion-label position="fixed">Username</ion-label>
            <ion-input clear-input></ion-input>
          </ion-row>
          <ion-row>
          <ion-label position="fixed">Password</ion-label>
          <ion-input type="password" clear-input></ion-input>
          </ion-row>
        </ion-item>
      </ion-grid>

    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent } from '@ionic/vue';
import ExploreContainer from '@/components/ExploreContainer.vue';

export default  {
  name: 'Tab1',
  components: { ExploreContainer, IonHeader, IonToolbar, IonTitle, IonContent, IonPage }
}
</script>

<style>
  .redline{
    border: 5px solid #781e1e;
    width: 100%;
    height: 5px;
    /*transform: rotate(-45deg);*/
    position: fixed;
    /*left: -15%;*/
  }
</style>