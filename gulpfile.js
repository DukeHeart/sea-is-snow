var gulp = require('gulp');
var watch = require('gulp-watch');
var livereload = require('gulp-livereload');
var shell = require('gulp-shell')

gulp.task('watch', function() {
  livereload.listen();
  gulp.watch('default/**.*').on('change', livereload.changed);
});

gulp.task('appserver', shell.task([
  '~/appengine/dev_appserver.py app.yaml'
]));

gulp.task('update', shell.task([
  'vulcanize -o default/build.html default/index.html --inline --csp --strip',
  '~/appengine/appcfg.py --oauth2 update app.yaml'
]));

gulp.task('dispatch', shell.task([
  '~/appengine/appcfg.py --oauth2 update_dispatch dispatch.yaml'
]));

gulp.task('index', shell.task([
  '~/appengine/appcfg.py --oauth2 update_indexes index.yaml'
]));

gulp.task('vacuum', shell.task([
  '~/appengine/appcfg.py --oauth2 vacuum_indexes index.yaml'
]));

gulp.task('rollback', shell.task([
  '~/appengine/appcfg.py --oauth2 rollback app.yaml'
]));

gulp.task('build', shell.task([
  'vulcanize -o default/build.html default/index.html --inline --csp --strip'
]));

gulp.task('default', ['watch', 'appserver']);

/*
var vulcanize = require('gulp-vulcanize');
var webserver = require('gulp-webserver');

gulp.task('build', function() {
  return gulp.src('../www/index.html')
    .pipe(vulcanize({dest: './', inline: true, strip: true }))
    .pipe(gulp.dest('./'));
 });

gulp.task('webserver', function() {
  gulp.src('../www')
    .pipe(webserver({
      livereload: true,
      directoryListing: true,
      open: true,
      fallback: 'index.html'
    }));
});
*/
