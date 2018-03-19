select "TRACKER----->",name from databases order by name;
select "FORM-------->",name from forms order by name;
select "DASHBOARD--->",name from dashboards order by name;
select "LIST-------->",name from item_lists order by name;
select "FIELD------->",name,type,ref_name,is_calc_field from fields order by name;
select "VALUE LIST-->",name,properties from value_lists order by name;
select "FORM COMP--->",forms.name,form_components.type from form_components,forms where forms.form_id=form_components.form_id order by forms.name,form_components.type;
select "DASH COMP--->",dashboards.name,dashboard_components.type from dashboard_components,dashboards where dashboards.dashboard_id=dashboard_components.dashboard_id order by dashboards.name,dashboard_components.type;