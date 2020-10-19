const sources = new Image();
sources.src = './assets/sources.png';

const canvas = document.querySelector('canvas');
const context = canvas.getContext('2d');

function hasCollision(flappyBird, element) {
  return flappyBird.y - flappyBird.height >= element.y;
}

const GetReady = {
  sourceX: 134,
  sourceY: 0,
  width: 174,
  height: 152,
  x: canvas.width / 2 - 174 / 2,
  y: 50,

  draw() {
    context.drawImage(
      sources,
      GetReady.sourceX,
      GetReady.sourceY,
      GetReady.width,
      GetReady.height,
      GetReady.x,
      GetReady.y,
      GetReady.width,
      GetReady.height
    );
  },
};

const FlappyBird = {
  sourceX: 0,
  sourceY: 0,
  width: 33,
  height: 24,
  x: 10,
  y: 50,
  gravity: 0.25,
  speed: 0,
  jumpValue: 4.6,

  jump() {
    FlappyBird.speed = -FlappyBird.jumpValue;
  },

  refresh() {
    if (hasCollision(FlappyBird, Ground)) {
      changeScreen(Screens.begin);
    }

    FlappyBird.speed += FlappyBird.gravity;
    FlappyBird.y += FlappyBird.speed;
  },

  draw() {
    context.drawImage(
      sources,
      FlappyBird.sourceX,
      FlappyBird.sourceY,
      FlappyBird.width,
      FlappyBird.height,
      FlappyBird.x,
      FlappyBird.y,
      FlappyBird.width,
      FlappyBird.height
    );
  },
};

const Ground = {
  sourceX: 0,
  sourceY: 610,
  width: 224,
  height: 112,
  x: 0,
  y: canvas.height - 112,

  draw() {
    context.drawImage(
      sources,
      Ground.sourceX,
      Ground.sourceY,
      Ground.width,
      Ground.height,
      Ground.x,
      Ground.y,
      Ground.width,
      Ground.height
    );

    context.drawImage(
      sources,
      Ground.sourceX,
      Ground.sourceY,
      Ground.width,
      Ground.height,
      Ground.x + Ground.width,
      Ground.y,
      Ground.width,
      Ground.height
    );
  },
};

const Background = {
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
      Background.sourceX,
      Background.sourceY,
      Background.width,
      Background.height,
      Background.x,
      Background.y,
      Background.width,
      Background.height
    );

    context.drawImage(
      sources,
      Background.sourceX,
      Background.sourceY,
      Background.width,
      Background.height,
      Background.x + Background.width,
      Background.y,
      Background.width,
      Background.height
    );
  },
};

let currentScreen = {};
function changeScreen(screen) {
  currentScreen = screen;
}

const Screens = {
  begin: {
    draw() {
      Background.draw();
      Ground.draw();
      FlappyBird.draw();
      GetReady.draw();
    },

    click() {
      FlappyBird.sourceX = 0;
      FlappyBird.sourceY = 0;
      FlappyBird.width = 33;
      FlappyBird.height = 24;
      FlappyBird.x = 10;
      FlappyBird.y = 50;
      FlappyBird.gravity = 0.25;
      FlappyBird.speed = 0;
      FlappyBird.jumpValue = 4.6;
      changeScreen(Screens.game);
    },

    refresh() {},
  },

  game: {
    draw() {
      Background.draw();
      Ground.draw();
      FlappyBird.draw();
    },

    click() {
      FlappyBird.jump();
    },

    refresh() {
      FlappyBird.refresh();
    },
  },
};

function run() {
  currentScreen.draw();
  currentScreen.refresh();

  requestAnimationFrame(run);
}

window.addEventListener('click', function () {
  if (currentScreen.click) {
    currentScreen.click();
  }
});

changeScreen(Screens.begin);
run();
