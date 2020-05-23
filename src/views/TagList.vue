<template>
    <div>
        <h2>
            {{ tag(this.$route.params.id).Name }}
        </h2>
    </div>
</template>

<script>
export default {
  computed: {
    tag () {
      return this.$store.state.videos.video.Tags.find(tag => tag.id === this.ID)
    }
  }
}
</script>
