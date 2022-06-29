# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_description_entity import AdminDescriptionEntity  # noqa: F401,E501
from flyteadmin.models.core_workflow_template import CoreWorkflowTemplate  # noqa: F401,E501


class AdminWorkflowSpec(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'template': 'CoreWorkflowTemplate',
        'sub_workflows': 'list[CoreWorkflowTemplate]',
        'description_entity': 'AdminDescriptionEntity'
    }

    attribute_map = {
        'template': 'template',
        'sub_workflows': 'sub_workflows',
        'description_entity': 'description_entity'
    }

    def __init__(self, template=None, sub_workflows=None, description_entity=None):  # noqa: E501
        """AdminWorkflowSpec - a model defined in Swagger"""  # noqa: E501

        self._template = None
        self._sub_workflows = None
        self._description_entity = None
        self.discriminator = None

        if template is not None:
            self.template = template
        if sub_workflows is not None:
            self.sub_workflows = sub_workflows
        if description_entity is not None:
            self.description_entity = description_entity

    @property
    def template(self):
        """Gets the template of this AdminWorkflowSpec.  # noqa: E501

        Template of the task that encapsulates all the metadata of the workflow.  # noqa: E501

        :return: The template of this AdminWorkflowSpec.  # noqa: E501
        :rtype: CoreWorkflowTemplate
        """
        return self._template

    @template.setter
    def template(self, template):
        """Sets the template of this AdminWorkflowSpec.

        Template of the task that encapsulates all the metadata of the workflow.  # noqa: E501

        :param template: The template of this AdminWorkflowSpec.  # noqa: E501
        :type: CoreWorkflowTemplate
        """

        self._template = template

    @property
    def sub_workflows(self):
        """Gets the sub_workflows of this AdminWorkflowSpec.  # noqa: E501

        Workflows that are embedded into other workflows need to be passed alongside the parent workflow to the propeller compiler (since the compiler doesn't have any knowledge of other workflows - ie, it doesn't reach out to Admin to see other registered workflows).  In fact, subworkflows do not even need to be registered.  # noqa: E501

        :return: The sub_workflows of this AdminWorkflowSpec.  # noqa: E501
        :rtype: list[CoreWorkflowTemplate]
        """
        return self._sub_workflows

    @sub_workflows.setter
    def sub_workflows(self, sub_workflows):
        """Sets the sub_workflows of this AdminWorkflowSpec.

        Workflows that are embedded into other workflows need to be passed alongside the parent workflow to the propeller compiler (since the compiler doesn't have any knowledge of other workflows - ie, it doesn't reach out to Admin to see other registered workflows).  In fact, subworkflows do not even need to be registered.  # noqa: E501

        :param sub_workflows: The sub_workflows of this AdminWorkflowSpec.  # noqa: E501
        :type: list[CoreWorkflowTemplate]
        """

        self._sub_workflows = sub_workflows

    @property
    def description_entity(self):
        """Gets the description_entity of this AdminWorkflowSpec.  # noqa: E501

        DescriptionEntity encapsulates all the detailed documentation for the workflow.  # noqa: E501

        :return: The description_entity of this AdminWorkflowSpec.  # noqa: E501
        :rtype: AdminDescriptionEntity
        """
        return self._description_entity

    @description_entity.setter
    def description_entity(self, description_entity):
        """Sets the description_entity of this AdminWorkflowSpec.

        DescriptionEntity encapsulates all the detailed documentation for the workflow.  # noqa: E501

        :param description_entity: The description_entity of this AdminWorkflowSpec.  # noqa: E501
        :type: AdminDescriptionEntity
        """

        self._description_entity = description_entity

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminWorkflowSpec, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminWorkflowSpec):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
