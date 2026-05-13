CREATE TABLE todo_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    color TEXT DEFAULT '#6b7280',
    description TEXT
);

CREATE INDEX idx_todo_categories_user_id ON todo_categories(user_id);

CREATE UNIQUE INDEX todo_categories_unique_name 
ON todo_categories(user_id, name);

CREATE TRIGGER set_updated_at_todo_categories
    BEFORE UPDATE ON todo_categories
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE todos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'draft',
    priority TEXT NOT NULL DEFAULT 'medium',
    due_date TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    parent_todo_id UUID REFERENCES todos,
    category_id UUID REFERENCES todo_categories ON DELETE SET NULL,
    metadata JSONB,
    sort_order SERIAL
);

CREATE TABLE todo_comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    todo_id UUID NOT NULL REFERENCES todos ON DELETE CASCADE,
    user_id TEXT NOT NULL,
    content TEXT NOT NULL
);

---- create above / drop below ----