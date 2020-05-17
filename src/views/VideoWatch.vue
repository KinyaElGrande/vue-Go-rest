<template>
    <div>
        <h1>{{video.Name}}</h1>
        <div v-html="video.Description"></div>
        <div v-for="Tag in video.Tags" :key="Tag.ID">
          {{ Tag.Name}}
        </div>
    </div>
</template>

<script>
export default {
  mounted () {
    this.$store.dispatch('loadVideos')
  },
  computed: {
    video () {
      return this.$store.state.videos.find(vid => vid.id === this.$route.params.ID)
    }
  }
}
</script>
