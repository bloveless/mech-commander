<script lang="ts">
  import { EventsOn } from "../wailsjs/runtime/runtime";

  let baseX = 17.5 * 0;
  let baseY = 17.5 * 0;

  let mechX = 17.5 * 15;
  let mechY = 17.5 * 15;

  let enemyMechX = 17.5 * 10;
  let enemyMechY = 17.5 * 10;

  let enemyBaseX = 17.5 * 39;
  let enemyBaseY = 17.5 * 39;

  EventsOn("game-tick", (data: any) => {
    mechX = data.x;
    mechY = data.y;
    console.debug("tick", data);
  });
</script>

<main>
  <div class="gameFrame">
    <div class="statsBar">
      <div class="playerStats">Player: 3</div>
      <div class="enemyStats">Enemy: 3</div>
    </div>
    <div class="gameBoard">
      <div class="base player" style="top: {baseX}px; left {baseY}px;" />
      <div class="mech player" style="top: {mechX}px; left: {mechY}px;" />
      <div class="base enemy" style="top: {enemyBaseX}px; left: {enemyBaseY}px;" />
      <div class="mech enemy" style="top: {enemyMechX}px; left: {enemyMechY}px;" />
    </div>
  </div>
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
  }

  .gameFrame {
    padding: 15px;
  }

  .statsBar {
    display: flex;
    justify-content: space-between;
    font-size: 2rem;
    color: black;
  }

  .gameBoard {
    position: relative;
    box-sizing: content-box;
    aspect-ratio: 1/1;
    height: calc(700px + 1px);
    width: calc(700px + 1px);

    background-image: repeating-linear-gradient(
        #ccc 0 1px,
        transparent 1px 100%
      ),
      repeating-linear-gradient(90deg, #ccc 0 1px, transparent 1px 100%);
    background-size: 2.5% 2.5%;
  }

  .mech {
    position: absolute;
    margin-top: 0.28%;
    margin-left: 0.28%;
    height: 2%;
    width: 2%;
    border-radius: 50%;
  }

  .mech.player {
    background-color: orange;
  }

  .mech.enemy {
    background-color: red;
  }

  .base {
    position: absolute;
    height: 2.5%;
    width: 2.5%;
  }

  .base.player {
    background-color: green;
  }

  .base.enemy {
    background-color: blue;
  }
</style>
