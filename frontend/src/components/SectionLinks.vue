<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://riot.im/app/#/room/#akashnet:matrix.org"
      target="_blank"
      title="Community Chat"
      subtitle="#akashnet")
    list-item(
      href="https://t.me/AkashNW"
      target="_blank"
      title="Telegram"
      subtitle="#AkashNW")
    list-item(
      href="https://github.com/ovrclk/akash"
      target="_blank"
      title="Source Code"
      subtitle="ovrclk/akash")
    list-item(
      href="https://akash.network/"
      target="_blank"
      title="Website"
      subtitle="akash.network")
</template>

<script>
import ListItem from "../components/NiListItem.vue";
export default {
  name: "faucet",
  components: {
    ListItem
  },
  data: () => ({}),
  methods: {}
};
</script>
