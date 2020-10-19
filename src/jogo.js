const sources = new Image();
sources.src = './assets/sources.png';

const canvas = document.querySelector('canvas');
const context = canvas.getContext('2d');

const flappyBird = {
  sourceX: 0,
  sourceY: 0,
  width: 33,
  height: 24,
  x: 10,
  y: 50,
  gravity: 0.25,
  speed: 0,

  refresh() {
    flappyBird.speed += flappyBird.gravity;
    flappyBird.y += 1;
  },

  draw() {
    context.drawImage(
      sources,
      flappyBird.sourceX,
      flappyBird.sourceY,
      flappyBird.width,
      flappyBird.height,
      flappyBird.x,
      flappyBird.y,
      flappyBird.width,
      flappyBird.height
    );
  },
};

const ground = {
  sourceX: 0,
  sourceY: 610,
  width: 224,
  height: 112,
  x: 0,
  y: canvas.height - 112,

  draw() {
    context.drawImage(
      sources,
      ground.sourceX,
      ground.sourceY,
      ground.width,
      ground.height,
      ground.x,
      ground.y,
      ground.width,
      ground.height
    );

    context.drawImage(
      sources,
      ground.sourceX,
      ground.sourceY,
      ground.width,
      ground.height,
      ground.x + ground.width,
      ground.y,
      ground.width,
      ground.height
    );
  },
};

const background = {
  sourceX: 390,
  sourceY: 0,
  width: 275,
  height: 204,
  x: 0,
  y: canvas.height - 204,

  draw() {
    context.fillStyle = '#70c5ce';
    context.fillRect(0, 0, canvas.width, canvas.height);

    context.drawImage(
      sources,
      background.sourceX,
      background.sourceY,
      background.width,
      background.height,
      background.x,
      background.y,
      background.width,
      background.height
    );

    context.drawImage(
      sources,
      background.sourceX,
      background.sourceY,
      background.width,
      background.height,
      background.x + background.width,
      background.y,
      background.width,
      background.height
    );
  },
};

function run() {
  flappyBird.refresh();
  background.draw();
  ground.draw();
  flappyBird.draw();

  requestAnimationFrame(run);
}

run();
