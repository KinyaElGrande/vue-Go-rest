<template>
    <div>
        <h1>{{video.name}}</h1>
        <div v-html="video.description"></div>
    </div>
</template>

<script>
export default {
  computed: {
    video () {
      return this.$store.state.videos.find(vid => vid.id === this.$route.params.id)
    }
  }
}
</script>
