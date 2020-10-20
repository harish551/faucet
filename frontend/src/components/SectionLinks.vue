<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://discord.gg/W8trcGV"
      target="_blank"
      title="Community Chat"
      subtitle="#cosmosnetwork")
    list-item(
      href="https://t.me/cosmosproject"
      target="_blank"
      title="Telegram"
      subtitle="#cosmosproject")
    list-item(
      href="https://github.com/cosmos/testnets"
      target="_blank"
      title="Testnet repo"
      subtitle="cosmos/testnets")
    list-item(
      href="https://cosmos.network/"
      target="_blank"
      title="Website"
      subtitle="cosmos.network")
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
