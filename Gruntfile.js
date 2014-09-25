'use strict';

module.exports = function (grunt) {

  grunt.initConfig({

    gae: {
      run: {
        action: 'run',
        options: {async:true, asyncOutput: true} //path
      },
      kill: {action: 'kill'},
      update: {action: 'update'},
      index: {action: 'update_indexes'},
      vacuum: {action: 'vacuum_indexes'},
      rollback: {action: 'rollback'}
    },

    watch: {
      options: {spawn: false},
      livereload: {
        options: {livereload: true},
        files: ['default/**']
      }
    },

    vulcanize: {
      default: {
        options: {
          inline: true,
          strip: true
        },
        files: {
          'default/build.html': 'default/index.html'
        }
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-vulcanize');
  grunt.loadNpmTasks('grunt-gae');
  grunt.registerTask('default', ['vulcanize','gae:run','watch']);
  
};
