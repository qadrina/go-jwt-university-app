var courseObject = {
    "Programming": {
        "C#": ["Aug 05-10 | 08.00 - 12.00",
                "Aug 05-10 | 08.00 - 12.00",
                "Aug 05-10 | 08.00 - 12.00"],
        "C++": ["Aug 06-10 | 08.00 - 12.00",
                "Aug 06-10 | 13.00 - 12.00",
                "Aug 06-10 | 08.00 - 12.00"],
        "Python": ["Aug 07-10 | 08.00 - 12.00",
                "Aug 07-10 | 13.00 - 12.00",
                "Aug 07-10 | 08.00 - 12.00"]    
    },
    "Microsoft Office": {
        "Word": ["Aug 10-15 | 08.00 - 12.00",
                "Aug 10-15 | 08.00 - 12.00",
                "Aug 10-15 | 08.00 - 12.00"],
        "Excel": ["Aug 12-17 | 08.00 - 12.00",
                "Aug 12-17 | 08.00 - 12.00",
                "Aug 05-10 | 08.00 - 12.00"],
        "PowerPoint": ["Aug 15-20 | 08.00 - 12.00",
                "Aug 15-20 | 08.00 - 12.00",
                "Aug 05-10 | 08.00 - 12.00"]
    },
    "Database": {
        "SQL Server": ["Aug 15-10 | 08.00 - 12.00",
                       "Aug 15-10 | 08.00 - 12.00",
                       "Aug 15-10 | 08.00 - 12.00"],
        "MySQL": ["Aug 20-10 | 08.00 - 12.00",
                  "Aug 20-10 | 13.00 - 12.00",
                  "Aug 05-10 | 08.00 - 12.00"],
        "Firebase": ["Aug 25-10 | 08.00 - 12.00",
                     "Aug 05-10 | 13.00 - 12.00",
                     "Aug 05-10 | 08.00 - 12.00"]
    }
  }

  window.onload = function() {
    var courseSel = document.getElementById("course");
    var subjectSel = document.getElementById("subject");
    var scheduleSel = document.getElementById("schedule");
    for (var x in courseObject) {
      courseSel.options[courseSel.options.length] = new Option(x, x);
    }
    courseSel.onchange = function() {
      //empty Chapters- and Topics- dropdowns
      scheduleSel.length = 1;
      subjectSel.length = 1;
      //display correct values
      for (var y in courseObject[this.value]) {
        subjectSel.options[subjectSel.options.length] = new Option(y, y);
      }
    }
    subjectSel.onchange = function() {
      //empty Chapters dropdown
      scheduleSel.length = 1;
      //display correct values
      var z = courseObject[courseSel.value][this.value];
      for (var i = 0; i < z.length; i++) {
        scheduleSel.options[scheduleSel.options.length] = new Option(z[i], z[i]);
      }
    }
  }