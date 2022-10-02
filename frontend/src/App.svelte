<script lang="ts">
  import { EventsOn } from '../wailsjs/runtime/runtime';

  let mechX = 10;
  let mechY = 10;

  EventsOn('game-tick', (data: any) => {
    mechX = data.x;
    mechY = data.y;
    console.debug('tick', data);
  });
</script>

<main>
  <div class="gameBoard">
    <div class="player" style="top: {mechX}px; left: {mechY}px;" />
    <div class="enemy" style="bottom: 10px; right: 10px;" />
  </div>
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
  }

  .gameBoard {
    position: relative;
    width: 100%;
    height: 100%;
  }

  .player {
    position: absolute;
    display: inline-block;
    background-color: orange;
    content: " ";
    height: 10px;
    width: 10px;
    border-radius: 50%;
  }

  .enemy {
    position: absolute;
    display: inline-block;
    background-color: red;
    content: " ";
    height: 10px;
    width: 10px;
    border-radius: 50%;
  }
</style>
