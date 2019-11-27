package projects

import "testing"

func init() {
	DevelopmentLogger()
	defer Logger.Sync()
}

var sample = `<?xml version="1.0" ?>
<qdbapi>
    <action>API_DoQuery</action>
    <errcode>0</errcode>
    <errtext>No error</errtext>
    <dbinfo>
        <name>Projects</name>
        <desc></desc>
    </dbinfo>
    <variables>
</variables>
    <chdbids>
</chdbids>
    <record rid="1287">
        <__p_>&lt;a title=&quot;Profile&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/113-note.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__p_>
        <__m_>&lt;a title=&quot;Microblog&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/115-book_open2.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__m_>
        <__l_>&lt;a title=&quot;Project Plan&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/721-task.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__l_>
        <__a_>&lt;a title=&quot;Analyses&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=15&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/145-powerpoint.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__a_>
        <__b_>&lt;a title=&quot;Budgets&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/443-dollars.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__b_>
        <ps>&lt;div&gt;&lt;img src=&quot;https://images.quickbase.com/si/16/240-triang_green.png&quot;&gt;&lt;/div&gt;</ps>
        <project_status>Active</project_status>
        <reason_closed>Contract/Work Was Completed</reason_closed>
        <partner>Bandujo</partner>
        <client>Kumon</client>
        <project_name>Kumon - GAC - 12/31/2019</project_name>
        <product>GAC</product>
        <____>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/012-options_1.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</____>
        <start_date>1546300800000</start_date>
        <end_date>1577750400000</end_date>
        <account_directors/>
        <managers>steve.kozma@befoundonline.com</managers>
        <analysts/>
        <jr__analysts/>
        <creatives/>
        <records>1287</records>
        <add_record/>
        <records2>1287</records2>
        <add_record2/>
        <related_client>206</related_client>
        <___of_budgets>17</___of_budgets>
        <___of_budgets_this_year>17</___of_budgets_this_year>
        <___of_contracts/>
        <___of_posts/>
        <___used_of_total_budget_hours/>
        <_____>&lt;a title=&quot;Setup&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/512-tools.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</_____>
        <__r_>&lt;a title=&quot;Roles&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/748-inheritance_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__r_>
        <__s_>&lt;a title=&quot;Schedule&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=29&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/813-tree_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__s_>
        <actual_spend_fees___total>0.00</actual_spend_fees___total>
        <add_contract/>
        <add_initiative>https://befoundonline.quickbase.com/db/bgqctxzys?a=API_GenAddRecordForm&amp;_fid_9=1287&amp;z=</add_initiative>
        <add_many_tasks>https://www.quickbase.com/db/bgpsdzvzs?a=q&amp;qid=23&amp;nv=1&amp;v0=</add_many_tasks>
        <add_monthly_budget>https://befoundonline.quickbase.com/db/bj22i7dry?a=API_GenAddRecordForm&amp;_fid_16=1287&amp;z=</add_monthly_budget>
        <add_post/>
        <add_task_type/>
        <added_to_ccbm_>No</added_to_ccbm_>
        <adjusted_budget_hours_this_month>10.00</adjusted_budget_hours_this_month>
        <archive>0</archive>
        <billable_fees_last_quarter>11700.00</billable_fees_last_quarter>
        <billable_fees_last_year/>
        <billable_fees_less_non_billable_initiatives_ptd>35100.00</billable_fees_less_non_billable_initiatives_ptd>
        <billable_fees_ptd>42900.00</billable_fees_ptd>
        <billable_fees_this_month>3900.00</billable_fees_this_month>
        <copy_project__with_tasks_>javascript:void(copyMasterDetailButtonHandler(&#039;&amp;copyFid=8&amp;destrid=0&amp;sourcerid=1287&#039;))</copy_project__with_tasks_>
        <billable_fees_this_quarter>11700.00</billable_fees_this_quarter>
        <billable_fees_this_year>46800.00</billable_fees_this_year>
        <billable_fees_total>46800.00</billable_fees_total>
        <billable_fees_ytd>42900.00</billable_fees_ytd>
        <billing_records>1287</billing_records>
        <budget_doc___project_plan_links/>
        <budgeted_hours_entire_project>190.00</budgeted_hours_entire_project>
        <budgeted_hours_last_quarter>40</budgeted_hours_last_quarter>
        <budgeted_hours_ptd>180.00</budgeted_hours_ptd>
        <budgeted_hours_this_month>10.00</budgeted_hours_this_month>
        <budgeted_hours_this_quarter>30</budgeted_hours_this_quarter>
        <budgeted_hours_this_year>190</budgeted_hours_this_year>
        <budgeted_hours_ytd>180.00</budgeted_hours_ytd>
        <ccbm_project_details>1287</ccbm_project_details>
        <client___exec_relationship_manager>dan.golden@befoundonline.com</client___exec_relationship_manager>
        <client_billable_>1</client_billable_>
        <client_form_view_icon___analyses>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=20&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/145-powerpoint.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___analyses>
        <client_form_view_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=23&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/113-note.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___profile>
        <client_form_view___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=23</client_form_view___profile>
        <client_form_view_icon___tasks>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=22&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/721-task.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___tasks>
        <client_full_name>Kumon North America</client_full_name>
        <client_logo/>
        <client_form_view___tasks>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=22</client_form_view___tasks>
        <contract_type>MRR (&gt; 3 months)</contract_type>
        <contracts>1287</contracts>
        <copy_project__without_tasks_>javascript:void(copyMasterDetailButtonHandler(&#039;&amp;relfids=31&amp;recurse=false&amp;copyFid=8&amp;destrid=0&amp;sourcerid=1287&#039;, &#039;bgpsa9a5i&#039;))</copy_project__without_tasks_>
        <creation_access>beth.spiegel@befoundonline.com;julia.ebner@befoundonline.com</creation_access>
        <client_form_view___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=20</client_form_view___analyses>
        <current_user>uris77@gmail.com</current_user>
        <date_created>1546453258143</date_created>
        <date_modified>1554228549842</date_modified>
        <effective_hourly_rate___ptd/>
        <form_view_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/113-note.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___profile>
        <elapsed_project_time>0.90</elapsed_project_time>
        <form_edit___ccbm>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=30</form_edit___ccbm>
        <form_view_icon___schedule>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=29&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/813-tree_view.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___schedule>
        <form_edit_icon___ccbm>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=30&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___ccbm>
        <form_view___ccbm>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=30</form_view___ccbm>
        <form_view_icon___ccbm>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=30&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/113-note.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___ccbm>
        <full_internal_access>0</full_internal_access>
        <furthest_budget_month_fom>1575158400000</furthest_budget_month_fom>
        <furthest_entered_budget_month_eom>1577750400000</furthest_entered_budget_month_eom>
        <hour_spend_run_rate/>
        <hours_available_entire_project/>
        <form_view_icon___roles>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/748-inheritance_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___roles>
        <form_view_icon___project_plan>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/721-task.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___project_plan>
        <hours_available_last_quarter/>
        <hours_available_ptd/>
        <hours_available_this_month/>
        <hours_available_this_quarter/>
        <hours_available_this_year/>
        <hours_available_ytd/>
        <form_view_icon___analyses>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=15&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/145-powerpoint.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___analyses>
        <hours_planned_this_month/>
        <hours_spent_this_month/>
        <form_view_icon___budgets>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/443-dollars.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___budgets>
        <initiatives>1287</initiatives>
        <internal_project_grade/>
        <form_view_icon___setup>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/512-tools.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___setup>
        <labor_budget_remaining/>
        <last_modified_by>beth.spiegel@befoundonline.com</last_modified_by>
        <microblog>1287</microblog>
        <form_view_icon___microblog>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/115-book_open2.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___microblog>
        <monthly_budgets>1287</monthly_budgets>
        <monthly_budgets2>1287</monthly_budgets2>
        <form_view_icon___field_decriptions>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/012-options_1.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___field_decriptions>
        <next_budget_alert____/>
        <form_view___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=14</form_view___profile>
        <non_labor_expenses_last_quarter>0.00</non_labor_expenses_last_quarter>
        <non_labor_expenses_this_month>0.00</non_labor_expenses_this_month>
        <non_labor_expenses_this_quarter>0.00</non_labor_expenses_this_quarter>
        <form_view___schedule>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=29</form_view___schedule>
        <non_labor_expenses_total>0.00</non_labor_expenses_total>
        <non_labor_expenses_ytd>0.00</non_labor_expenses_ytd>
        <form_view___roles>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=24</form_view___roles>
        <ops_level_access>accounting@befoundonline.com;beth.spiegel@befoundonline.com;dan.golden@befoundonline.com;julia.ebner@befoundonline.com;steve.krull@befoundonline.com</ops_level_access>
        <form_view___project_plan>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=17</form_view___project_plan>
        <partner_full_name>Bandujo</partner_full_name>
        <partner_logo/>
        <form_view___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=15</form_view___analyses>
        <profit_margin_entire_project/>
        <profit_margin_last_quarter/>
        <profit_margin_ptd/>
        <profit_margin_this_month/>
        <form_view___budgets>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=16</form_view___budgets>
        <profit_margin_this_quarter/>
        <profit_outlook_entire_project/>
        <profit_outlook_last_quarter/>
        <form_view___setup>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=18</form_view___setup>
        <profit_outlook_this_month/>
        <profit_outlook_this_quarter/>
        <profitability_entire_project/>
        <profitability_last_quarter/>
        <profitability_this_month/>
        <profitability_this_quarter/>
        <project>Kumon - GAC - 12/31/2019</project>
        <form_view___microblog>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=26</form_view___microblog>
        <project_id>1287</project_id>
        <form_view___field_descriptions>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1287&amp;dfid=25</form_view___field_descriptions>
        <project_id___billable>0</project_id___billable>
        <project_leads>steve.kozma@befoundonline.com</project_leads>
        <form_edit_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___profile>
        <project_length>31449600000</project_length>
        <project_lookup_display>Kumon - GAC - 12/31/2019</project_lookup_display>
        <project_support/>
        <form_edit_icon___microblog>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___microblog>
        <proposed_ad/>
        <form_edit_icon___project_plan>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___project_plan>
        <proposed_analyst/>
        <proposed_manager/>
        <record_owner>beth.spiegel@befoundonline.com</record_owner>
        <renewal_status/>
        <restricted>0</restricted>
        <service_description/>
        <form_edit_icon___roles>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___roles>
        <form_edit_icon___budgets>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___budgets>
        <form_edit_icon___setup>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___setup>
        <form_edit_icon___field_descriptions>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___field_descriptions>
        <form_edit___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=14</form_edit___profile>
        <staffing_usage_next_month/>
        <form_edit___microblog>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=26</form_edit___microblog>
        <start_date_fom>1546300800000</start_date_fom>
        <task_types>1287</task_types>
        <form_edit___roles>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=24</form_edit___roles>
        <form_edit___project_plan>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=17</form_edit___project_plan>
        <form_edit___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=15</form_edit___analyses>
        <team__name/>
        <form_edit___budgets>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=16</form_edit___budgets>
        <total_labor_budget___used/>
        <total_us_budget_hours_available/>
        <vertical>Education - new</vertical>
        <form_edit___setup>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=18</form_edit___setup>
        <form_edit___field_descriptions>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1287&amp;dfid=25</form_edit___field_descriptions>
        <add_budget>https://befoundonline.quickbase.com/db/bg44v3kd9?a=API_GenAddRecordForm&amp;_fid_6=1287&amp;dfid=10&amp;z=</add_budget>
        <access_condition_level>Unrestricted</access_condition_level>
        <contract_status>Signed</contract_status>
        <associated_contract_ids/>
        <updated_in_sf>0</updated_in_sf>
        <project_id___initiative/>
        <project_id___initiative_name/>
        <deal_type>Renewal</deal_type>
        <prebill_media>0</prebill_media>
        <prebill_management>0</prebill_management>
        <hs_product>Analytics</hs_product>
        <prebill_tech>0</prebill_tech>
        <mgmt_fee_calc/>
        <client_deposit>0</client_deposit>
        <bfo_pays_media>0</bfo_pays_media>
        <media_deposit_amount/>
        <media_deposit_notes/>
        <media_deposit_invoice_nbr/>
        <records3>1287</records3>
        <add_record3>https://befoundonline.quickbase.com/db/bpyeznxmd?a=API_GenAddRecordForm&amp;_fid_13=1287&amp;z=</add_record3>
        <nonstandard_billing>0</nonstandard_billing>
        <media_deposit_returned>0</media_deposit_returned>
        <update_id>1554228549842</update_id>
    </record>
    <record rid="1288">
        <__p_>&lt;a title=&quot;Profile&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/113-note.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__p_>
        <__m_>&lt;a title=&quot;Microblog&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/115-book_open2.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__m_>
        <__l_>&lt;a title=&quot;Project Plan&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/721-task.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__l_>
        <__a_>&lt;a title=&quot;Analyses&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=15&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/145-powerpoint.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__a_>
        <__b_>&lt;a title=&quot;Budgets&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/443-dollars.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__b_>
        <ps>&lt;div&gt;&lt;img src=&quot;https://images.quickbase.com/si/16/240-triang_green.png&quot;&gt;&lt;/div&gt;</ps>
        <project_status>Active</project_status>
        <reason_closed/>
        <partner>BFO Direct</partner>
        <client>European Wax Center (Regional)</client>
        <project_name>EWC Corporate (r) John Mok - Wyckoff - SEM</project_name>
        <product>SEM</product>
        <____>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/012-options_1.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</____>
        <start_date>1546300800000</start_date>
        <end_date/>
        <account_directors>josh.barwa@befoundonline.com</account_directors>
        <managers>ryan.mccullough@befoundonline.com</managers>
        <analysts>raquel.kink@befoundonline.com</analysts>
        <jr__analysts/>
        <creatives/>
        <records>1288</records>
        <add_record/>
        <records2>1288</records2>
        <add_record2/>
        <related_client>323</related_client>
        <___of_budgets>13</___of_budgets>
        <___of_budgets_this_year>13</___of_budgets_this_year>
        <___of_contracts/>
        <___of_posts/>
        <___used_of_total_budget_hours/>
        <_____>&lt;a title=&quot;Setup&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/512-tools.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</_____>
        <__r_>&lt;a title=&quot;Roles&quot; href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/748-inheritance_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__r_>
        <__s_>&lt;a title=&quot;Schedule&quot;href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=29&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/24/813-tree_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</__s_>
        <actual_spend_fees___total>2340.00</actual_spend_fees___total>
        <add_contract/>
        <add_initiative>https://befoundonline.quickbase.com/db/bgqctxzys?a=API_GenAddRecordForm&amp;_fid_9=1288&amp;z=</add_initiative>
        <add_many_tasks>https://www.quickbase.com/db/bgpsdzvzs?a=q&amp;qid=23&amp;nv=1&amp;v0=</add_many_tasks>
        <add_monthly_budget>https://befoundonline.quickbase.com/db/bj22i7dry?a=API_GenAddRecordForm&amp;_fid_16=1288&amp;z=</add_monthly_budget>
        <add_post/>
        <add_task_type/>
        <added_to_ccbm_>Yes</added_to_ccbm_>
        <adjusted_budget_hours_this_month>1.00</adjusted_budget_hours_this_month>
        <archive>0</archive>
        <billable_fees_last_quarter>585.00</billable_fees_last_quarter>
        <billable_fees_last_year/>
        <billable_fees_less_non_billable_initiatives_ptd>2535.00</billable_fees_less_non_billable_initiatives_ptd>
        <billable_fees_ptd>2535.00</billable_fees_ptd>
        <billable_fees_this_month>195.00</billable_fees_this_month>
        <copy_project__with_tasks_>javascript:void(copyMasterDetailButtonHandler(&#039;&amp;copyFid=8&amp;destrid=0&amp;sourcerid=1288&#039;))</copy_project__with_tasks_>
        <billable_fees_this_quarter>585.00</billable_fees_this_quarter>
        <billable_fees_this_year>2730.00</billable_fees_this_year>
        <billable_fees_total>2730.00</billable_fees_total>
        <billable_fees_ytd>2535.00</billable_fees_ytd>
        <billing_records>1288</billing_records>
        <budget_doc___project_plan_links/>
        <budgeted_hours_entire_project>14.00</budgeted_hours_entire_project>
        <budgeted_hours_last_quarter>3</budgeted_hours_last_quarter>
        <budgeted_hours_ptd>13.00</budgeted_hours_ptd>
        <budgeted_hours_this_month>1.00</budgeted_hours_this_month>
        <budgeted_hours_this_quarter>3</budgeted_hours_this_quarter>
        <budgeted_hours_this_year>14</budgeted_hours_this_year>
        <budgeted_hours_ytd>13.00</budgeted_hours_ytd>
        <ccbm_project_details>1288</ccbm_project_details>
        <client___exec_relationship_manager/>
        <client_billable_>1</client_billable_>
        <client_form_view_icon___analyses>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=20&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/145-powerpoint.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___analyses>
        <client_form_view_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=23&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/113-note.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___profile>
        <client_form_view___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=23</client_form_view___profile>
        <client_form_view_icon___tasks>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=22&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/721-task.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</client_form_view_icon___tasks>
        <client_full_name>European Wax Center (Regional)</client_full_name>
        <client_logo/>
        <client_form_view___tasks>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=22</client_form_view___tasks>
        <contract_type>MRR (&gt; 3 months)</contract_type>
        <contracts>1288</contracts>
        <copy_project__without_tasks_>javascript:void(copyMasterDetailButtonHandler(&#039;&amp;relfids=31&amp;recurse=false&amp;copyFid=8&amp;destrid=0&amp;sourcerid=1288&#039;, &#039;bgpsa9a5i&#039;))</copy_project__without_tasks_>
        <creation_access>julia.ebner@befoundonline.com</creation_access>
        <client_form_view___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=20</client_form_view___analyses>
        <current_user>uris77@gmail.com</current_user>
        <date_created>1546461880652</date_created>
        <date_modified>1565737959826</date_modified>
        <effective_hourly_rate___ptd/>
        <form_view_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/113-note.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___profile>
        <elapsed_project_time/>
        <form_edit___ccbm>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=30</form_edit___ccbm>
        <form_view_icon___schedule>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=29&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/813-tree_view.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___schedule>
        <form_edit_icon___ccbm>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=30&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___ccbm>
        <form_view___ccbm>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=30</form_view___ccbm>
        <form_view_icon___ccbm>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=30&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/113-note.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___ccbm>
        <full_internal_access>0</full_internal_access>
        <furthest_budget_month_fom>1575158400000</furthest_budget_month_fom>
        <furthest_entered_budget_month_eom>1577750400000</furthest_entered_budget_month_eom>
        <hour_spend_run_rate/>
        <hours_available_entire_project/>
        <form_view_icon___roles>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/748-inheritance_view.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___roles>
        <form_view_icon___project_plan>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/721-task.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___project_plan>
        <hours_available_last_quarter/>
        <hours_available_ptd/>
        <hours_available_this_month/>
        <hours_available_this_quarter/>
        <hours_available_this_year/>
        <hours_available_ytd/>
        <form_view_icon___analyses>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=15&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/145-powerpoint.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___analyses>
        <hours_planned_this_month/>
        <hours_spent_this_month/>
        <form_view_icon___budgets>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/443-dollars.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___budgets>
        <initiatives>1288</initiatives>
        <internal_project_grade>A</internal_project_grade>
        <form_view_icon___setup>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/512-tools.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___setup>
        <labor_budget_remaining/>
        <last_modified_by>beth.spiegel@befoundonline.com</last_modified_by>
        <microblog>1288</microblog>
        <form_view_icon___microblog>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/115-book_open2.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___microblog>
        <monthly_budgets>1288</monthly_budgets>
        <monthly_budgets2>1288</monthly_budgets2>
        <form_view_icon___field_decriptions>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/128/012-options_1.png&quot; width=&quot;64&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_view_icon___field_decriptions>
        <next_budget_alert____/>
        <form_view___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=14</form_view___profile>
        <non_labor_expenses_last_quarter>0.00</non_labor_expenses_last_quarter>
        <non_labor_expenses_this_month>0.00</non_labor_expenses_this_month>
        <non_labor_expenses_this_quarter>0.00</non_labor_expenses_this_quarter>
        <form_view___schedule>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=29</form_view___schedule>
        <non_labor_expenses_total>0.00</non_labor_expenses_total>
        <non_labor_expenses_ytd>0.00</non_labor_expenses_ytd>
        <form_view___roles>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=24</form_view___roles>
        <ops_level_access>accounting@befoundonline.com;beth.spiegel@befoundonline.com;dan.golden@befoundonline.com;julia.ebner@befoundonline.com;steve.krull@befoundonline.com</ops_level_access>
        <form_view___project_plan>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=17</form_view___project_plan>
        <partner_full_name>Be Found Online</partner_full_name>
        <partner_logo/>
        <form_view___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=15</form_view___analyses>
        <profit_margin_entire_project/>
        <profit_margin_last_quarter/>
        <profit_margin_ptd/>
        <profit_margin_this_month/>
        <form_view___budgets>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=16</form_view___budgets>
        <profit_margin_this_quarter/>
        <profit_outlook_entire_project/>
        <profit_outlook_last_quarter/>
        <form_view___setup>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=18</form_view___setup>
        <profit_outlook_this_month/>
        <profit_outlook_this_quarter/>
        <profitability_entire_project/>
        <profitability_last_quarter/>
        <profitability_this_month/>
        <profitability_this_quarter/>
        <project>EWC Corporate (r) John Mok - Wyckoff - SEM</project>
        <form_view___microblog>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=26</form_view___microblog>
        <project_id>1288</project_id>
        <form_view___field_descriptions>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=dr&amp;rid=1288&amp;dfid=25</form_view___field_descriptions>
        <project_id___billable>0</project_id___billable>
        <project_leads>josh.barwa@befoundonline.com;raquel.kink@befoundonline.com;ryan.mccullough@befoundonline.com</project_leads>
        <form_edit_icon___profile>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=14&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___profile>
        <project_length/>
        <project_lookup_display>EWC Corporate (r) John Mok - Wyckoff - SEM</project_lookup_display>
        <project_support>raquel.kink@befoundonline.com</project_support>
        <form_edit_icon___microblog>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=26&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___microblog>
        <proposed_ad/>
        <form_edit_icon___project_plan>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=17&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___project_plan>
        <proposed_analyst>ryan.mccullough@befoundonline.com</proposed_analyst>
        <proposed_manager>ryan.mccullough@befoundonline.com</proposed_manager>
        <record_owner>mark@befoundonline.com</record_owner>
        <renewal_status/>
        <restricted>0</restricted>
        <service_description/>
        <form_edit_icon___roles>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=24&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___roles>
        <form_edit_icon___budgets>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=16&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___budgets>
        <form_edit_icon___setup>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=18&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___setup>
        <form_edit_icon___field_descriptions>&lt;a href=&quot;https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=25&quot;&gt;&lt;img src=&quot;https://images.quickbase.com/si/48/005-pen_0.png&quot;&gt;&lt;/img&gt;&lt;/a&gt;</form_edit_icon___field_descriptions>
        <form_edit___profile>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=14</form_edit___profile>
        <staffing_usage_next_month/>
        <form_edit___microblog>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=26</form_edit___microblog>
        <start_date_fom>1546300800000</start_date_fom>
        <task_types>1288</task_types>
        <form_edit___roles>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=24</form_edit___roles>
        <form_edit___project_plan>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=17</form_edit___project_plan>
        <form_edit___analyses>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=15</form_edit___analyses>
        <team__name/>
        <form_edit___budgets>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=16</form_edit___budgets>
        <total_labor_budget___used/>
        <total_us_budget_hours_available/>
        <vertical/>
        <form_edit___setup>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=18</form_edit___setup>
        <form_edit___field_descriptions>https://befoundonline.quickbase.com/db/bgpsa9a5i?a=er&amp;rid=1288&amp;dfid=25</form_edit___field_descriptions>
        <add_budget>https://befoundonline.quickbase.com/db/bg44v3kd9?a=API_GenAddRecordForm&amp;_fid_6=1288&amp;dfid=10&amp;z=</add_budget>
        <access_condition_level>Unrestricted</access_condition_level>
        <contract_status>Signed</contract_status>
        <associated_contract_ids/>
        <updated_in_sf>1</updated_in_sf>
        <project_id___initiative/>
        <project_id___initiative_name/>
        <deal_type>New Business</deal_type>
        <prebill_media>0</prebill_media>
        <prebill_management>0</prebill_management>
        <hs_product>Paid Media - Paid Search, Social, &amp; Remarketing</hs_product>
        <prebill_tech>0</prebill_tech>
        <mgmt_fee_calc>Net</mgmt_fee_calc>
        <client_deposit>0</client_deposit>
        <bfo_pays_media>1</bfo_pays_media>
        <media_deposit_amount/>
        <media_deposit_notes/>
        <media_deposit_invoice_nbr/>
        <records3>1288</records3>
        <add_record3>https://befoundonline.quickbase.com/db/bpyeznxmd?a=API_GenAddRecordForm&amp;_fid_13=1288&amp;z=</add_record3>
        <nonstandard_billing>0</nonstandard_billing>
        <media_deposit_returned>0</media_deposit_returned>
        <update_id>1565737959826</update_id>
    </record>
</qdbapi>
`

func TestUnmarshallProject(t *testing.T) {
	doc := UnmarshallProject(sample)
	projects := doc.Projects
	if len(projects) != 2 {
		t.Error("there should be 2 projects")
	}
}
