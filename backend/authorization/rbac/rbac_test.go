package rbac

import (
	"context"
	"errors"
	"testing"

	"github.com/echlebek/sensu-lite/backend/authorization"
	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/testing/mockstore"
	"github.com/echlebek/sensu-lite/types"
	"github.com/stretchr/testify/mock"
)

func TestAuthorize(t *testing.T) {
	type storeFunc func(*mockstore.MockStore)
	var nilClusterRoleBindings []*types.ClusterRoleBinding
	var nilRoleBindings []*types.RoleBinding
	tests := []struct {
		name      string
		attrs     *authorization.Attributes
		storeFunc storeFunc
		want      bool
		wantErr   bool
	}{
		{
			name:  "no bindings",
			attrs: &authorization.Attributes{Namespace: "acme"},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)
				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilRoleBindings, nil)
			},
			want: false,
		},
		{
			name: "ClusterRoleBindings store err",
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "no matching ClusterRoleBinding",
			attrs: &authorization.Attributes{
				Namespace: "acme",
				User: types.User{
					Username: "foo",
				},
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.ClusterRoleBinding{&types.ClusterRoleBinding{
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "bar"},
						},
					}}, nil)
				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilRoleBindings, nil)
			},
			want: false,
		},
		{
			name: "GetClusterRole store err",
			attrs: &authorization.Attributes{
				User: types.User{
					Username: "foo",
				},
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.ClusterRoleBinding{&types.ClusterRoleBinding{
						RoleRef: types.RoleRef{
							Type: "ClusterRole",
							Name: "admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetClusterRole", mock.AnythingOfType("*context.emptyCtx"), "admin", mock.Anything).
					Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "matching ClusterRoleBinding",
			attrs: &authorization.Attributes{
				Verb:         "create",
				Resource:     "checks",
				ResourceName: "check-cpu",
				User: types.User{
					Username: "foo",
				},
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.ClusterRoleBinding{&types.ClusterRoleBinding{
						RoleRef: types.RoleRef{
							Type: "ClusterRole",
							Name: "admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetClusterRole", mock.AnythingOfType("*context.emptyCtx"), "admin", mock.Anything).
					Return(&types.ClusterRole{Rules: []types.Rule{
						types.Rule{
							Verbs:         []string{"create"},
							Resources:     []string{"checks"},
							ResourceNames: []string{"check-cpu"},
						},
					}}, nil)
			},
			want: true,
		},
		{
			name:  "RoleBindings store err",
			attrs: &authorization.Attributes{Namespace: "acme"},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)
				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilRoleBindings, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "no matching RoleBindings",
			attrs: &authorization.Attributes{
				Namespace: "acme",
				User: types.User{
					Username: "foo",
				},
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)
				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.RoleBinding{&types.RoleBinding{
						RoleRef: types.RoleRef{
							Type: "Role",
							Name: "admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetRole", mock.Anything, "admin").
					Return(nil, nil)
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "GetRole store err",
			attrs: &authorization.Attributes{
				Namespace: "acme",
				User: types.User{
					Username: "foo",
				},
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)
				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.RoleBinding{&types.RoleBinding{
						RoleRef: types.RoleRef{
							Type: "Role",
							Name: "admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetRole", mock.Anything, "admin").
					Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "matching RoleBinding",
			attrs: &authorization.Attributes{
				Namespace: "acme",
				User: types.User{
					Username: "foo",
				},
				Verb:         "create",
				Resource:     "checks",
				ResourceName: "check-cpu",
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)

				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.RoleBinding{&types.RoleBinding{
						RoleRef: types.RoleRef{
							Type: "Role",
							Name: "admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetRole", mock.Anything, "admin").
					Return(&types.Role{Rules: []types.Rule{
						types.Rule{
							Verbs:         []string{"create"},
							Resources:     []string{"checks"},
							ResourceNames: []string{"check-cpu"},
						},
					}}, nil)
			},
			want: true,
		},
		{
			name: "role bindings do not match cluster width resource request",
			attrs: &authorization.Attributes{
				User: types.User{
					Username: "foo",
				},
				Verb:     "list",
				Resource: "users",
			},
			storeFunc: func(s *mockstore.MockStore) {
				s.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return(nilClusterRoleBindings, nil)

				s.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
					Return([]*types.RoleBinding{&types.RoleBinding{
						RoleRef: types.RoleRef{
							Type: "ClusterRole",
							Name: "cluster-admin",
						},
						Subjects: []types.Subject{
							types.Subject{Type: types.UserType, Name: "foo"},
						},
					}}, nil)
				s.On("GetClusterRole", mock.AnythingOfType("*context.emptyCtx"), "cluster-admin", mock.Anything).
					Return(&types.ClusterRole{Rules: []types.Rule{
						types.Rule{
							Verbs:     []string{"*"},
							Resources: []string{"*"},
						},
					}}, nil)
			},
			want: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := &mockstore.MockStore{}
			a := &Authorizer{
				Store: store,
			}
			tc.storeFunc(store)

			got, err := a.Authorize(context.Background(), tc.attrs)
			if (err != nil) != tc.wantErr {
				t.Errorf("Authorizer.Authorize() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("Authorizer.Authorize() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMatchesUser(t *testing.T) {
	tests := []struct {
		name     string
		user     types.User
		subjects []types.Subject
		want     bool
	}{
		{
			name: "not matching",
			user: types.User{Username: "foo"},
			subjects: []types.Subject{
				types.Subject{Type: types.UserType, Name: "bar"},
				types.Subject{Type: types.GroupType, Name: "foo"},
			},
			want: false,
		},
		{
			name: "matching via username",
			user: types.User{Username: "foo"},
			subjects: []types.Subject{
				types.Subject{Type: types.UserType, Name: "bar"},
				types.Subject{Type: types.UserType, Name: "foo"},
			},
			want: true,
		},
		{
			name: "matching via group",
			user: types.User{Username: "foo", Groups: []string{"acme"}},
			subjects: []types.Subject{
				types.Subject{Type: types.GroupType, Name: "acme"},
			},
			want: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := matchesUser(tc.user, tc.subjects); got != tc.want {
				t.Errorf("matchesUser() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRuleAllows(t *testing.T) {
	tests := []struct {
		name  string
		attrs *authorization.Attributes
		rule  types.Rule
		want  bool
	}{
		{
			name: "verb does not match",
			attrs: &authorization.Attributes{
				Verb: "create",
			},
			rule: types.Rule{
				Verbs: []string{"get"},
			},
			want: false,
		},
		{
			name: "resource does not match",
			attrs: &authorization.Attributes{
				Verb:     "create",
				Resource: "events",
			},
			rule: types.Rule{
				Verbs:     []string{"create"},
				Resources: []string{"checks", "handlers"},
			},
			want: false,
		},
		{
			name: "resource name does not match",
			attrs: &authorization.Attributes{
				Verb:         "create",
				Resource:     "checks",
				ResourceName: "check-cpu",
			},
			rule: types.Rule{
				Verbs:         []string{"create"},
				Resources:     []string{"checks"},
				ResourceNames: []string{"check-mem"},
			},
			want: false,
		},
		{
			name: "matches",
			attrs: &authorization.Attributes{
				Verb:         "create",
				Resource:     "checks",
				ResourceName: "check-cpu",
			},
			rule: types.Rule{
				Verbs:         []string{"create"},
				Resources:     []string{"checks"},
				ResourceNames: []string{"check-cpu"},
			},
			want: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := ruleAllows(tc.attrs, tc.rule); got != tc.want {
				t.Errorf("ruleAllows() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestVisitRulesFor(t *testing.T) {
	attrs := &authorization.Attributes{
		Namespace: "acme",
		User: types.User{
			Username: "foo",
		},
		Verb:         "create,delete",
		Resource:     "checks",
		ResourceName: "check-cpu",
	}
	stor := &mockstore.MockStore{}
	a := &Authorizer{
		Store: stor,
	}
	stor.On("ListClusterRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
		Return([]*types.ClusterRoleBinding{&types.ClusterRoleBinding{
			RoleRef: types.RoleRef{
				Type: "ClusterRole",
				Name: "admin",
			},
			Subjects: []types.Subject{
				types.Subject{Type: types.UserType, Name: "foo"},
			},
		}}, nil)

	stor.On("ListRoleBindings", mock.AnythingOfType("*context.emptyCtx"), &store.SelectionPredicate{}).
		Return([]*types.RoleBinding{&types.RoleBinding{
			RoleRef: types.RoleRef{
				Type: "Role",
				Name: "admin",
			},
			Subjects: []types.Subject{
				types.Subject{Type: types.UserType, Name: "foo"},
			},
		}}, nil)
	stor.On("GetRole", mock.Anything, "admin").
		Return(&types.Role{Rules: []types.Rule{
			types.Rule{
				Verbs:         []string{"create"},
				Resources:     []string{"checks"},
				ResourceNames: []string{"check-cpu"},
			},
		}}, nil)
	stor.On("GetClusterRole", mock.AnythingOfType("*context.emptyCtx"), "admin", mock.Anything).
		Return(&types.ClusterRole{Rules: []types.Rule{
			types.Rule{
				Verbs:         []string{"delete"},
				Resources:     []string{"checks"},
				ResourceNames: []string{"check-cpu"},
			},
		}}, nil)

	var rules []types.Rule

	a.VisitRulesFor(context.Background(), attrs, func(binding RoleBinding, rule types.Rule, err error) bool {
		if err != nil {
			t.Fatal(err)
			return false
		}
		rules = append(rules, rule)
		return true
	})

	if got, want := len(rules), 2; got != want {
		t.Fatalf("wrong number of rules: got %d, want %d", got, want)
	}
}
