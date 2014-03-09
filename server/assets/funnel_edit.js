// Renders a funnel's steps to the form.
function render(funnel) {
    var root = $(".steps")
    root.empty();

    for (var i=0; i<funnel.steps.length; i++) {
        appendStep(funnel.steps[i]);
    }
}

// Appends a step form control to the end of the steps section.
function appendStep(step) {
    var index = $(".steps .form-group").length;
    var group = $('<div class="form-group">');
    var label = $('<label for="name">Step #' + (index+1) + '</label>');
    var select = $('<select class="form-control">')
        .attr("id", "step[" + index + "].condition")
        .attr("name", "step[" + index + "].condition");

    for (var i=0; i<resources.length; i++) {
      var value = 'resource == "' + resources[i] + '"';
      console.log("?", value, step.condition, step.condition === value);
      var option = $("<option>")
          .val(value)
          .text(resources[i])
          .attr("selected", step.condition === value);

      select.append(option);
    }

    group.append(label).append(select);
    $(".steps").append(group);
}

$(document).on("click", ".add-step", function() {
    $(".steps", appendStep({condition: ""}));
});
