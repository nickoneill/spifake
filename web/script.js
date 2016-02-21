var numberLights = 749;

document.addEventListener( 'DOMContentLoaded', function () {
  console.log("loaded");

  setupCanvas();

  var spisocket = new WebSocket("ws://localhost:5480/spi");
  // spisocket.binaryType = "arraybuffer";
  // var exampleSocket = new WebSocket("ws://www.example.com/socketserver", "protocolOne");

  spisocket.onopen = function (event) {
    console.log("opened");
  };

  spisocket.onmessage = function (event) {
    // console.log("got message",event.data);

    var reader = new FileReader();
    reader.addEventListener("loadend", function() {
      var dv = new DataView(reader.result);
      // console.log("buffer",dv.byteLength,dv.getUint8(3+((82+20)*4)));
      drawLights(dv);
    });
    reader.readAsArrayBuffer(event.data);
  }
}, false );

function getCanvasContext() {
  var c = document.getElementById("drawable");
  var ctx = c.getContext("2d");

  return ctx;
}

function setupCanvas() {
  var ctx = getCanvasContext();

  ctx.fillStyle = "#000";
  ctx.fillRect(0,0,600,600);
}

function rgbaForLightIndex(lights, index, offset) {
  // opacity bump! give an extra 50% because pixels != leds
  var a = ((lights.getUint8(4+((offset+parseInt(index))*4)) - 224) / 31) + 0.5;
  var b = lights.getUint8(4+((offset+parseInt(index))*4)+1);
  var g = lights.getUint8(4+((offset+parseInt(index))*4)+2);
  var r = lights.getUint8(4+((offset+parseInt(index))*4)+3);

  // if (offset+(index*4) == 497) {
  //   var uint = lights.getUint8(3+offset+(index*4) - 224);
  //   var a = (lights.getUint8(3+offset+(index*4)) - 224) / 31;
  //   var b = lights.getUint8(3+offset+(index*4)+1);
  //   var g = lights.getUint8(3+offset+(index*4)+2);
  //   var r = lights.getUint8(3+offset+(index*4)+3);
  //   // var alpha = (lights.getUint8(4+offset+(index*4)) - 224);
  //   console.log("c",a,b,g,r);
  // }

  var rgba = "rgba(" + r +"," + g + "," + b + "," + a + ")";
  // console.log("pixel",rgba);
  return rgba;
}

function drawLights(lights) {
  var ctx = getCanvasContext();
  ctx.clearRect(0,0,800,900);

  ctx.fillStyle = "#000";
  ctx.fillRect(0,0,800,900);

  var drawOffsetX = 150;
  var drawOffsetY = 50;

  console.log("draw");

  // draw right side
  // 82 -> 429 : 347

  for (y in _.range(0, 347)) {
      ctx.fillStyle = rgbaForLightIndex(lights,y,81);
      ctx.fillRect(drawOffsetX + 400, drawOffsetY + (y*2), 2, 2);
  }

  // draw front
  // 430 -> 612 : 182
  for (x in _.range(0, 182)) {
      ctx.fillStyle = rgbaForLightIndex(lights,x,429);
      ctx.fillRect(drawOffsetX + (400 - x*2), drawOffsetY + (347*2), 2, 2);
  }

  // draw left
  // 613 -> 748 : 135
  for (y in _.range(0, 135)) {
      ctx.fillStyle = rgbaForLightIndex(lights,y,613);
      ctx.fillRect(drawOffsetX + (400 - (182*2)), drawOffsetY + (347 - y)*2, 2, 2);
  }
}

function backingScale(context) {
    if ('devicePixelRatio' in window) {
        if (window.devicePixelRatio > 1) {
            return window.devicePixelRatio;
        }
    }
    return 1;
}
