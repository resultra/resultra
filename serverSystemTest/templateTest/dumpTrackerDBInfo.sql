select "TRACKER----->",name from databases order by name;
select "FORM-------->",name from forms order by name;
select "TABLE VIEW-->",name from table_views order by name;
select "TABLE COL--->",table_views.name,table_view_columns.type from table_views,table_view_columns where table_view_columns.table_id=table_views.table_id order by table_view_columns.type;
select "DASHBOARD--->",name from dashboards order by name;
select "LIST-------->",name from item_lists order by name;
select "NEW ITEM LINK---->",name from form_links order by name;
select "FIELD------->",name,type,ref_name,is_calc_field from fields order by name;
select "VALUE LIST-->",name,properties from value_lists order by name;
select "FORM COMP--->",forms.name,form_components.type from form_components,forms where forms.form_id=form_components.form_id order by forms.name,form_components.type;
select "DASH COMP--->",dashboards.name,dashboard_components.type from dashboard_components,dashboards where dashboards.dashboard_id=dashboard_components.dashboard_id order by dashboards.name,dashboard_components.type;