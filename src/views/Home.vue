<template>
  <div>
      <div v-for="video in videos" :key="video.ID">
        <router-link :to="{name: 'video-watch', params: {id: video.ID}}">
          <h3>{{ video.Name }}</h3>
        </router-link>
      </div>
  </div>
</template>

<script>

export default {
  name: 'Home',
  computed: {
    videos () { return this.$store.state.videos }
  }
}
</script>
