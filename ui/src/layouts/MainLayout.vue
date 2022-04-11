<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Sonrai
        </q-toolbar-title>

        <div>Sonrai v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item-label
          header
        >
          Navigation
        </q-item-label>

        <EssentialLink
          v-for="link in essentialLinks"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'
import EssentialLink from 'components/EssentialLink.vue'

const linksList = [
  {
    title: 'Producers',
    caption: '3 of 7 producers running',
    icon: 'code',
    link: 'https://quasar.dev'
  },
  {
    title: 'Data Records',
    caption: '5 data types defined',
    icon: 'note_add',
    link: 'https://github.com/quasarframework'
  },
  {
    title: 'Monitors',
    caption: '0 alarms',
    icon: 'leaderboard',
    link: 'https://chat.quasar.dev'
  },
  {
    title: 'Syncs',
    caption: '1 sync',
    icon: 'record_voice_over',
    link: 'https://forum.quasar.dev'
  }
]

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink
  },

  setup () {
    const leftDrawerOpen = ref(false)

    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
})
</script>
