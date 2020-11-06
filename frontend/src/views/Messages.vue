<template>
  <div>
    <layout
      top-nav
      logo="https://i.ibb.co/5GFxj30/handcuffs-2.png"
      :router="true"
      :side-nav="sideNav"
      toggle="full"
      class="messages"
    >
      <div class="container" slot="content">
        <div class="row">
          <div class="col-md-12">
            <div class="row mb-4">
            <header class="col-md-5 mb-3">
              <h1 style="font-size: 4rem;" class="mt-5">Messages</h1>
            </header>
            <img class="col-md-6" src="../assets/messages.png">
            </div>

            <table
              class="my-3 table table-striped table-compact table-hover table-primary"
            >
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Reply status</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>
                    xx
                  </td>
                  <td>xx</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </layout>
  </div>
</template>

<script>
import { Layout } from "rbc-wm-framework-vuejs/dist/wm/components";
import sideNav from "../../sidenav.JSON";
export default {
  name: "Messages",
  components: {
    "layout": Layout
  },
  data() {
    return {
      dynamicComponent: "Discover",
      sideNav: sideNav
    };
  }
};
</script>
<style>
.messages {
  padding-left: 5%;
  padding-top: 0%;
}
</style>
