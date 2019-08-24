<template>
  <div class="flex flex-col" id="app">
    <div class="mb-12" id="nav"></div>
    <router-view />
    <nav class="nav">
      <router-link to="/onboarding">
        <span class="nav__link">Onboarding</span>
      </router-link>
      <router-link to="/">
        <span class="nav__link">Home</span>
      </router-link>
    </nav>
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
