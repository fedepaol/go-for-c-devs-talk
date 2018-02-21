var result = tasks.map(function (task) {
                    return (task.duration / 60);
                }).filter(function (duration) {
                    return duration >= 2;
                });
