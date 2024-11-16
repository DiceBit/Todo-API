create table if not exists Tasks
(
    Id          integer primary key autoincrement,

    Title       varchar(255) not null,
    Description text,
    DueDate     varchar(255),

    Overdue     boolean      not null default false,

    Completed   boolean               default false
);

create index if not exists DueDate_idx on Tasks (DueDate)