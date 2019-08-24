<template>
  <div class="home">
    <WeatherDisplay />
    <InstagramImages hashtag="gornergrat" />
    Home!
  </div>
</template>

<script>
import InstagramImages from '../components/InstagramImages.vue';
import WeatherDisplay from '../components/WeatherDisplay.vue';
export default {
  name: 'home',
  components: { WeatherDisplay, InstagramImages }
};
</script>

<style lang="postcss">
body {
  @apply bg-GornerGratGold;
}
</style>
