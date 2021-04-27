package db

import (
	"fmt"
	"testing"
)

const UsersXml1 = `
<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<select id="getUserRoles"   resultType="cloud.ecosphere.yy.base.service.model.Role"  parameterType="java.lang.String">
			select distinct r.* from
			(
				select a.* from BASE_ROLE a,BASE_USER_ROLE b
				where a.ID_ = b.ROLE_ID_
				<if test="id!=null and id != ''">
				AND b.USER_ID_ = #{id}
				</if>
				UNION 
				select a.* from BASE_ROLE a,BASE_ROLE_POSITION b,BASE_USER_POSITION c
				where a.ID_ = b.ROLE_ID_ and b.POSITION_ID_ = c.POSITION_ID_ 
				<if test="id!=null and id != ''">
					and c.USER_ID_ = #{id}
				</if>
			) r
	</select>
	
	<select id="getUserPermissions"   resultType="cloud.ecosphere.yy.base.service.model.Permission"  parameterType="java.lang.String">
			select distinct r.* from
				(
				select  distinct c.* from BASE_USER_ROLE a,BASE_ROLE_PERMISSION b,BASE_PERMISSION c
				where  a.ROLE_ID_ = b.ROLE_ID_ and b.PERMISSION_ID_ =c.ID_
				<if test="id!=null and id != ''">
				 AND a.USER_ID_ = #{id}
				 </if>
				UNION
				select distinct d.* from BASE_USER_POSITION a,BASE_ROLE_POSITION b,BASE_ROLE_PERMISSION c,BASE_PERMISSION d
				where a.POSITION_ID_ = b.POSITION_ID_ and b.ROLE_ID_ = c.ROLE_ID_ and c.PERMISSION_ID_ = d.ID_
				<if test="id!=null and id != ''">
				AND a.USER_ID_ = #{id}  
				</if>
			) r
	</select>
	
	
	
	<select id="findByRoleIdAndCompanyIds"  resultType="cloud.ecosphere.yy.base.service.model.User">
		select distinct a.* from BASE_USER a,BASE_USER_ROLE b 
		where a.ID_  = b.USER_ID_ and b.ROLE_ID_=#{id} and a.STATUS_ > 0
		<if test="cids!=null and cids.size>0">
	        AND a.COMPANY_ID_  in 
	        <foreach collection="cids" item="itemId" open="(" separator="," close=")">
	        	#{itemId}
	        </foreach>
        </if>
	</select>
	
	<select id="findByPositionIdAndCompanyIds"  resultType="cloud.ecosphere.yy.base.service.model.User">
			select distinct a.* from BASE_USER a,BASE_USER_POSITION b 
			where a.ID_  = b.USER_ID_ and b.POSITION_ID_=#{id} and a.STATUS_ > 0 
			<if test="cids!=null and cids.size>0">
		        AND a.COMPANY_ID_  in 
		        <foreach collection="cids" item="itemId" open="(" separator="," close=")">
		        	#{itemId}
		        </foreach>
	        </if>
	</select>
	
	<select id="findUserByPositionAndDepartment"  resultType="java.lang.String">
		select distinct a.id_ from base_user a
			inner join base_user_role b
			inner join base_role c
			where a.id_ = b.user_id_ and b.role_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ in
			<foreach collection="roleIds" item="itemId" open="(" separator="," close=")">
		        	#{itemId}
		    </foreach>
			and a.department_id_ = #{departmentId}
			union
			select distinct a.id_ from base_user a
			inner join base_user_position b
			inner join base_position c
			where a.id_ = b.user_id_ and b.position_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ = #{posId}
			and a.department_id_ = #{departmentId}
	</select>
	
	<select id="findUserByPositionAndCompany"  resultType="java.lang.String">
		select distinct a.id_ from base_user a
			inner join base_user_role b
			inner join base_role c
			where a.id_ = b.user_id_ and b.role_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ in 
			<foreach collection="roleIds" item="itemId" open="(" separator="," close=")">
		        	#{itemId}
		    </foreach>
			and a.company_id_ = #{companyId}
			union
			select distinct a.id_ from base_user a
			inner join base_user_position b
			inner join base_position c
			where a.id_ = b.user_id_ and b.position_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ = #{posId}
			and a.company_id_ = #{companyId}
	</select>
	
	<select id="findUserByPosition"  resultType="java.lang.String">
		select distinct a.id_ from base_user a
			inner join base_user_role b
			inner join base_role c
			where a.id_ = b.user_id_ and b.role_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ in 
			<foreach collection="roleIds" item="itemId" open="(" separator="," close=")">
		        	#{itemId}
		    </foreach>
			union
			select distinct a.id_ from base_user a
			inner join base_user_position b
			inner join base_position c
			where a.id_ = b.user_id_ and b.position_id_ = c.id_
			and a.status_ >=1 and c.status_ = 1
			and c.id_ = #{posId}
	</select>
	
	<select id="findUserByTeamIds"  resultType="cloud.ecosphere.yy.base.service.model.User">
		select 
		a.*
		 from base_user a,base_user_team b
		where a.ID_ = b.USER_ID_
		and b.TEAM_ID_ in
			<foreach collection="teamIds" item="itemId" open="(" separator="," close=")">
		        	#{itemId}
		    </foreach>
		and a.STATUS_ > 0
	</select>

	<select id="findDeputyUserByPositionIdAndCompanyIds" resultType="java.lang.String">
		select distinct a.USER_ID_
		from base_user_owner a,base_user_owner_position b
		where a.ID_ = b.OWNER_ID_
		and b.POSITION_ID_ = #{posId} and a.TYPE_ = 0
		and a.COMPANY_ID_ in
		<foreach collection="cids" item="itemId" open="(" separator="," close=")">
	        	#{itemId}
	    </foreach>
	</select>



	<select id="getAssignUserById"  resultType="cloud.ecosphere.yy.base.service.model.User"
	  parameterType="cloud.ecosphere.yy.base.service.vo.BpmAssignVo">
		  SELECT distinct
			a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			FROM BASE_USER a
			where a.ID_ IN (
			<foreach collection="userIds"  item="item"  index="index" separator="," >
				#{item}
			 </foreach>
			)
			and a.STATUS_ > 0
	</select>


	<select id="getAssignUserByRole"  resultType="cloud.ecosphere.yy.base.service.model.User"
	  parameterType="cloud.ecosphere.yy.base.service.vo.BpmAssignVo">
		  select
			distinct a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			from BASE_USER a ,BASE_USER_ROLE b,BASE_ROLE c
			where a.ID_ = b.USER_ID_ and b.ROLE_ID_ = c.ID_
			and c.ID_ in (
			<foreach collection="roleIds"  item="item"  index="index" separator="," >
			    #{item}
			</foreach>
			)
			and a.STATUS_ > 0
			<if test='companyId!=null and companyId!=""'>
			and a.COMPANY_ID_ = #{companyId}
			</if>
			<if test='departmentId!=null and departmentId!=""'>
			and a.DEPARTMENT_ID_ = #{departmentId}
			</if>
			union
			select
			distinct a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			from BASE_USER a,base_user_position b ,base_position c,base_role_position d
			where a.ID_ = b.USER_ID_ and b.POSITION_ID_ = c.ID_ and c.ID_ = d.POSITION_ID_
			and c.STATUS_ = 1
			and a.STATUS_ > 0
			and d.ROLE_ID_ in (
			<foreach collection="roleIds"  item="item"  index="index" separator="," >
			    #{item}
			</foreach>
			)
			<if test='companyId!=null and companyId!=""'>
			and a.COMPANY_ID_ = #{companyId}
			</if>
			<if test='departmentId!=null and departmentId!=""'>
			and a.DEPARTMENT_ID_ = #{departmentId}
			</if>
	</select>

	<select id="getAssignUserByPosition"  resultType="cloud.ecosphere.yy.base.service.model.User"
	  parameterType="cloud.ecosphere.yy.base.service.vo.BpmAssignVo">
		  SELECT distinct
			a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			FROM BASE_USER a
			left outer join BASE_USER_POSITION b on b.USER_ID_ = a.ID_
			left outer join BASE_POSITION c on c.ID_ = b.POSITION_ID_
			where c.ID_ IN (
				<foreach collection="positionIds"  item="item"  index="index" separator="," >
					#{item}
				 </foreach>
			)
			<if test='companyId!=null and companyId!=""'>
			and a.COMPANY_ID_ = #{companyId}
			</if>
			<if test='departmentId!=null and departmentId!=""'>
			and a.DEPARTMENT_ID_ = #{departmentId}
			</if>
			and a.STATUS_ > 0
	</select>

	<select id="getAssignSlaveUserByRole"  resultType="cloud.ecosphere.yy.base.service.model.User"
	  parameterType="cloud.ecosphere.yy.base.service.vo.BpmAssignVo">
		  select
			distinct a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			from BASE_USER a,base_user_owner b,base_user_owner_position c,base_position d,base_role_position e
			where a.ID_ = b.USER_ID_ and b.ID_ = c.OWNER_ID_ and c.POSITION_ID_ = d.ID_  and d.ID_ = e.POSITION_ID_
			and d.STATUS_ = 1
			and a.STATUS_ > 0
			and e.ROLE_ID_ in (
			<foreach collection="roleIds"  item="item"  index="index" separator="," >
			    #{item}
			</foreach>
			)
			<if test='companyId!=null and companyId!=""'>
			and b.COMPANY_ID_ = #{companyId}
			</if>
			<if test='departmentId!=null and departmentId!=""'>
			and b.DEPARTMENT_ID_ = #{departmentId}
			</if>
	</select>

	<select id="getAssignSlaveUserByPosition"  resultType="cloud.ecosphere.yy.base.service.model.User"
	  parameterType="cloud.ecosphere.yy.base.service.vo.BpmAssignVo">
		  select
			distinct a.ID_ id,a.USERNAME_ username,a.LOGINNAME_ loginname
			from BASE_USER a,base_user_owner b,base_user_owner_position c
			where a.ID_ = b.USER_ID_ and b.ID_ = c.OWNER_ID_
			and a.STATUS_ > 0
			and c.POSITION_ID_ in (
			<foreach collection="positionIds"  item="item"  index="index" separator="," >
				#{item}
			 </foreach>
			)
			<if test='companyId!=null and companyId!=""'>
			and b.COMPANY_ID_ = #{companyId}
			</if>
			<if test='departmentId!=null and departmentId!=""'>
			and b.DEPARTMENT_ID_ = #{departmentId}
			</if>
	</select>

	<sql id="pageImport">
		<if test='username!=null and username!=""'>
			and r.USERNAME_ like #{username}
		</if>
		<if test='loginname!=null and loginname!=""'>
			and r.LOGINNAME_ like #{loginname}
		</if>
		<if test='departmentId!=null and departmentId!=""'>
			and (
				r.DEPARTMENT_ID_ = #{departmentId}
				or a.DEPARTMENT_ID_ = #{departmentId}
			)
		</if>
		<if test='companyId!=null and companyId!=""'>
			and (r.COMPANY_ID_ = #{companyId}
				or a.COMPANY_ID_ = #{companyId}
			)
		</if>
		<if test='status!=null'>
			AND r.STATUS_ >= #{status}
		</if>
		<if test='companyIdIn!=null and companyIdIn.size>0'>
		   AND (
			 r.COMPANY_ID_ in (
				<foreach collection="companyIdIn"  item="item"  index="index" separator="," >
					#{item}
				 </foreach>
			)
		    or
		    	a.COMPANY_ID_ in (
				<foreach collection="companyIdIn"  item="item"  index="index" separator="," >
					#{item}
				 </foreach>
			)
		)
		</if>
		<if test='departmentIdIn!=null and departmentIdIn.size>0'>
			AND (
				r.DEPARTMENT_ID_ in (
				<foreach collection="departmentIdIn"  item="item"  index="index" separator="," >
					#{item}
				 </foreach>
				)
				or a.DEPARTMENT_ID_ in (
				<foreach collection="departmentIdIn"  item="item"  index="index" separator="," >
					#{item}
				 </foreach>
				)
			)
		</if>
		<choose>
			<when test="addressbook!=null and addressbook==1">
				and r.COMPANY_ID_ in (
					<foreach collection="userInfo.relationCompanyId"  item="item"  index="index" separator="," >
							#{item}
					</foreach>
				)
			</when>
			<when test="relationcompany==true">
				and r.COMPANY_ID_ in (
					<foreach collection="userInfo.relationCompanyId"  item="item"  index="index" separator="," >
							#{item}
					</foreach>
				)
			</when>
			<when test="filialcompany==true">
				and r.COMPANY_ID_ in (
					<foreach collection="userInfo.filialCompanyId"  item="item"  index="index" separator="," >
							#{item}
					</foreach>
				)
			</when>
			<otherwise>
				and r.COMPANY_ID_ = #{userInfo.companyId}
			</otherwise>
		</choose>

	</sql>

	<select id="page1" resultType="cloud.ecosphere.yy.base.service.vo.UserVo"
	  parameterType="cloud.ecosphere.yy.base.service.form.UserForm">
		
		LEFT OUTER JOIN base_user_owner a on a.USER_ID_ = r.ID_
		LEFT OUTER JOIN BASE_DEPARTMENT r1 on r1.ID_=r.DEPARTMENT_ID_
		LEFT OUTER JOIN BASE_COMPANY r2 on r2.ID_=r.COMPANY_ID_
		LEFT OUTER JOIN BASE_DICTIONARY r3 on r3.VALUE_=r.STATUS_ AND r3.GROUPS_='system_platform_status'
		WHERE 1 = 1
		<include refid="pageImport"></include>
		ORDER BY r.CREATE_TIME_ desc,r.DEPARTMENT_ID_ desc,r.COMPANY_ID_ desc,r.STATUS_ desc
		limit #{firstResult},#{pageSize}
	</select>

	<select id="pageCount1" resultType="java.lang.Long"
	  parameterType="cloud.ecosphere.yy.base.service.form.UserForm">
		select count(1) from (
		SELECT distinct r.ID_
		FROM BASE_USER r
		LEFT OUTER JOIN base_user_owner a on a.USER_ID_ = r.ID_
		WHERE 1 = 1
		<include refid="pageImport"></include>
		) a1
	</select>
</mapper>

`

func TestName(t *testing.T) {
	ff := GetMapperFactory().Parse(nil, UsersXml1)
	fmt.Println(ff)
}
