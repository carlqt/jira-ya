<template>
  <div class="container">
      <div class="col-sm-3">
          <div class="panel panel-default">
              <div class="panel-heading clearfix">
                  <div class="col-xs-10">
                      <div class="animate-bg"></div>
                  </div>
              </div>
              <div class="panel-body">
                  <div class="col-sm-5">
                      <p class="animate-bg"><br><br><br></p>
                  </div>
                  <div class="col-sm-7">
                      <p class="animate-bg"></p>
                      <p class="animate-bg col-sm-11"></p>
                      <p class="animate-bg"></p>
                  </div>
              </div>
          </div>
          <div class="panel panel-default">
              <div class="panel-heading clearfix">
                  <div class="col-xs-10">
                      <div class="animate-bg"></div>
                  </div>
              </div>
              <div class="panel-body">
                  <div class="col-sm-5">
                      <p class="animate-bg"><br><br><br></p>
                  </div>
                  <div class="col-sm-7">
                      <p class="animate-bg"></p>
                      <p class="animate-bg col-sm-11"></p>
                      <p class="animate-bg"></p>
                  </div>
              </div>
          </div>
      </div>
     
      <div class="col-sm-9">
          <div class="panel panel-default">
              <div class="panel-heading clearfix">
                  <div class="col-xs-10">
                      <div class="animate-bg"></div>
                  </div>
              </div>
              <div class="panel-body">
                  <div class="clearfix">
                      <div class="col-sm-3">
                          <p class="animate-bg"><br><br><br><br><br><br></p>
                      </div>
                      <div class="col-sm-9">
                          <p class="animate-bg col-sm-11"></p>
                          <p class="animate-bg col-sm-9"></p>
                          <p class="animate-bg col-sm-10"></p>
                          <p class="animate-bg col-sm-9"></p>
                          <p class="animate-bg"></p>
                      </div>
                  </div>

                  <br>

                  <div class="col-sm-6">
                      <p class="animate-bg col-sm-11"></p>
                      <p class="animate-bg col-sm-9"></p>
                      <p class="animate-bg col-sm-10"></p>
                      <p class="animate-bg col-sm-9"></p>
                      <p class="animate-bg"></p>
                  </div>
                  <div class="col-sm-6">
                      <p class="animate-bg"><br><br><br><br><br><br></p>
                  </div>
              </div>
          </div>
      </div>
  </div>
</template>

<script>
export default {
  name: 'PageLoader'
}
</script>

<style>
@keyframes placeHolderShimmer {
	0% {
		background-position: -400px 0
	}
	100% {
		background-position: 400px 0
	}
}

.animate-bg {
	animation-duration: 1s;
	animation-fill-mode: forwards;
	animation-iteration-count: infinite;
	animation-name: placeHolderShimmer;
	animation-timing-function: linear;
	background: #f6f7f8;
	background: linear-gradient(to right, #eeeeee 8%, #dddddd 18%, #eeeeee 33%);
	background-size: 800px 104px;
	min-height: 12px;
	display: block;
	clear: both;
}
</style>
